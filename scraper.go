package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
)

var apiBaseUrl = "http://pubs.vmware.com/vsphere-55/topic/com.vmware.wssdk.apiref.doc"

var mosIndex = apiBaseUrl + "/index-mo_types.html"
var doIndex = apiBaseUrl + "/index-do_types.html"
var enumIndex = apiBaseUrl + "/index-e_types.html"
var faultIndex = apiBaseUrl + "/index-faults.html"

type Property struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	Description        string `json:"description"`
	Optional           bool   `json:"optional"`
	RequiredPrivileges string `json:"required_privileges"`
}

type Parameter struct {
	Name               string `json:"name"`
	Type               string `json:"type"`
	Description        string `json:"description"`
	Optional           bool   `json:"optional"`
	RequiredPrivileges string `json:"required_privileges"`
}

type Value struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Method struct {
	Name               string       `json:"name"`
	Description        string       `json:"description"`
	RequiredPrivileges []string     `json:"required_privileges"`
	Since              string       `json:"since"`
	Parameters         []*Parameter `json:"parameters"`
	ReturnValue        Value        `json:"return_value"`
	Faults             []Value      `json:"faults"`
	Events             []Value      `json:"events"`
}

type Object struct {
	Name        string      `json:"name"`
	Extends     string      `json:"extends"`
	Since       string      `json:"since"`
	Description string      `json:"description"`
	Properties  []*Property `json:"properties"`
	Methods     []*Method   `json:"methods"`
	Constants   []Constant  `json:"constants"`
}

type Constant struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

var objDescRegex *regexp.Regexp

func init() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	/**
	* HTML is not well formed so we need this regexp to extract
	* Objects descriptions.
	 */
	objDescRegex = regexp.MustCompile(`</h2>((?s).+?)<p class="table-title">`)

	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	log.SetPrefix("--> ")
}

func main() {
	channel := make(chan *Object)
	closeChannelAt := 0
	totalObjects := 0
	var objs []*Object

	go func() {
		closeChannelAt += scrape(mosIndex, channel)
	}()

	go func() {
		closeChannelAt += scrape(doIndex, channel)
	}()

	go func() {
		closeChannelAt += scrape(enumIndex, channel)
	}()

	go func() {
		closeChannelAt += scrape(faultIndex, channel)
	}()

	for obj := range channel {
		totalObjects++

		log.Println("type: " + obj.Name)
		objs = append(objs, obj)

		if totalObjects == closeChannelAt {
			break
		}
	}

	data, err := json.MarshalIndent(objs, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("./api.json", data, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("\n")
	log.Printf("Close channel at: %d\n", closeChannelAt)
	log.Printf("Total objects extracted: %d\n", totalObjects)
	log.Println("vSphere API definition was generated in ./api.json")
	log.Println("Done 💩")
}

func scrape(url string, channel chan *Object) int {
	d, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln(err)
	}

	handler := func(i int, sel *goquery.Selection) {
		href, _ := sel.Attr("href")
		name, _ := sel.Attr("title")

		go scrapeObject(href, name, channel)
	}

	ae := d.Find("#AE > nobr > a")
	fj := d.Find("#FJ > nobr > a")
	ko := d.Find("#KO > nobr > a")
	pt := d.Find("#PT > nobr > a")
	uz := d.Find("#UZ > nobr > a")

	closeChannelAt := ae.Length() + fj.Length() + ko.Length() + pt.Length() + uz.Length()

	ae.Each(handler)
	fj.Each(handler)
	ko.Each(handler)
	pt.Each(handler)
	uz.Each(handler)

	return closeChannelAt
}

func scrapeObject(path, name string, channel chan *Object) {
	d, err := goquery.NewDocument(apiBaseUrl + "/" + path)
	if err != nil {
		log.Fatalln(err)
	}

	obj := &Object{Name: name}

	d.Find("body > dl > dt").Each(func(i int, sel *goquery.Selection) {
		dtText := strings.TrimSpace(sel.Text())
		ddText := strings.TrimSpace(sel.Next().Text())
		if dtText == "Extends" {
			obj.Extends = ddText
		} else if dtText == "Since" {
			obj.Since = ddText
		}
	})

	html, err := d.Find("body").Html()
	if err != nil {
		log.Fatalln(err)
	}
	obj.Description = getObjectDesc(html)

	//log.Println("type: " + name)
	//log.Println(obj.Description)

	d.Find(`body > table`).Each(func(i int, sel *goquery.Selection) {
		prev := strings.TrimSpace(sel.Prev().Text())

		if prev == "Properties" {
			//obj.Properties = make([]*Property, sel.Find(`tbody > tr`).Length()-1)

			sel.Find("tbody > tr").Each(func(j int, sel2 *goquery.Selection) {
				if sel2.HasClass("r1") || sel2.HasClass("r0") {
					p := &Property{}
					pname := sel2.Find("td:nth-child(1)")

					p.Name = strings.TrimSpace(pname.Find("strong").Text())

					pname.Find("span").Each(func(k int, sel3 *goquery.Selection) {
						spanValue := strings.TrimSpace(sel3.Text())
						if spanValue == "*" {
							p.Optional = true
						} else if spanValue == "P" {
							p.RequiredPrivileges, _ = sel3.Attr("title")
						}
					})

					type_ := sel2.Find("td:nth-child(2) > a")
					if type_.Length() == 2 {
						//Instead of getting ManagedReferenceObjecte it gets
						//the real type
						p.Type = strings.TrimSpace(type_.Last().Text())
					} else if type_.Length() == 1 {
						p.Type = strings.TrimSpace(type_.Text())
					} else {
						//Gets primitive types such as: xsd:string, xsd:long, etc
						p.Type = strings.TrimSpace(sel2.Find("td:nth-child(2)").Text())
					}

					p.Description = strings.TrimSpace(sel2.Find("td:nth-child(3)").Text())

					if p.Name != "" {
						obj.Properties = append(obj.Properties, p)
						//log.Printf("%#v\n", p)
					}
				}
			})
		} else if prev == "Methods" {
			methods := sel.Find("tbody > tr:nth-child(2) > td > a")
			//obj.Methods = make([]*Method, methods.Length())

			methods.Each(func(i int, sel2 *goquery.Selection) {
				mname := sel2.Text()
				m := &Method{}
				m.Name = mname

				//log.Println("method: " + mname)

				m.Description = getMethodDesc(mname, html)

				d.Find("h1").Each(func(j int, sel3 *goquery.Selection) {
					if strings.TrimSpace(sel3.Text()) != mname {
						return
					}

					//Parameters
					sel4 := sel3.NextFilteredUntil(`p[class="table-title"]`, `table[cellspacing="0"]`).Next()
					//m.Parameters = make([]*Parameter, sel4.Find(`tbody > tr`).Length()-1)

					sel4.Find("tbody > tr").Each(func(k int, sel5 *goquery.Selection) {
						p := &Parameter{}
						pname := sel5.Find("td:nth-child(1)")

						p.Name = strings.TrimSpace(pname.Find("strong").Text())

						pname.Find("span").Each(func(k int, sel6 *goquery.Selection) {
							spanValue := strings.TrimSpace(sel6.Text())
							if spanValue == "*" {
								p.Optional = true
							} else if spanValue == "P" {
								p.RequiredPrivileges, _ = sel6.Attr("title")
							}
						})

						p.Type = strings.TrimSpace(sel5.Find("td:nth-child(2) > a").Text())
						p.Description = strings.TrimSpace(sel5.Find("td:nth-child(3)").Text())

						if p.Name != "" {
							m.Parameters = append(m.Parameters, p)
						}
					})

					//Return Value
					sel5 := sel4.NextFilteredUntil(`p[class="table-title"]`, `table[cellspacing="0"]`).Next()
					sel5.Find("tbody > tr").Each(func(k int, sel6 *goquery.Selection) {
						v := Value{}

						type_ := sel6.Find("td:nth-child(1) > a")
						if type_.Length() == 2 {
							//Instead of getting ManagedReferenceObjecte it gets
							//the real type
							v.Type = strings.TrimSpace(type_.Last().Text())
						} else if type_.Length() == 1 {
							v.Type = strings.TrimSpace(type_.Text())
						} else {
							//Gets primitive types such as: xsd:string, xsd:long, etc
							v.Type = strings.TrimSpace(sel6.Find("td:nth-child(1)").Text())
						}

						v.Description = sel6.Find("td:nth-child(2)").Text()
						if v.Type != "" {
							m.ReturnValue = v
							//log.Printf("%#v\n", v)
						}
					})

					//Faults
					sel6 := sel5.NextFilteredUntil(`p[class="table-title"]`, `table[cellspacing="0"]`).Next()
					//m.Faults = make([]Value, sel6.Find(`tbody > tr`).Length()-1)

					sel6.Find("tbody > tr").Each(func(k int, sel7 *goquery.Selection) {
						f := Value{}

						type_ := sel7.Find("td:nth-child(1) > a")
						if type_.Length() == 2 {
							//Instead of getting ManagedReferenceObjecte it gets
							//the real type
							f.Type = strings.TrimSpace(type_.Last().Text())
						} else if type_.Length() == 1 {
							f.Type = strings.TrimSpace(type_.Text())
						} else {
							//Gets primitive types such as: xsd:string, xsd:long, etc
							f.Type = strings.TrimSpace(sel7.Find("td:nth-child(1)").Text())
						}

						f.Description = sel7.Find("td:nth-child(2)").Text()
						if f.Type != "" {
							m.Faults = append(m.Faults, f)
							//log.Printf("%#v\n", f)
						}
					})
				})
				obj.Methods = append(obj.Methods, m)
			})
		} else if prev == "Enum Constants" {
			//obj.Constants = make([]Constant, sel.Find(`tbody > tr`).Length()-1)

			sel.Find("tbody > tr").Each(func(j int, sel2 *goquery.Selection) {
				c := Constant{}

				c.Name = sel2.Find("td:nth-child(1)").Text()
				c.Description = strings.TrimSpace(sel2.Find("td:nth-child(2)").Text())
				if c.Name != "" {
					obj.Constants = append(obj.Constants, c)
					//log.Printf("%#v\n", c)
				}
			})
		}
	})

	channel <- obj
}

func getObjectDesc(html string) string {
	matches := objDescRegex.FindStringSubmatch(html)
	if len(matches) == 2 {
		return strings.TrimSpace(matches[1])
	}

	return ""
}

func getMethodDesc(name, html string) string {
	methodDescRegex, err := regexp.Compile(`<h1>` + name + `</h1>((?s).+?)<p class="table-title">`)
	if err != nil {
		return ""
	}

	matches := methodDescRegex.FindStringSubmatch(html)
	if len(matches) == 2 {
		return strings.TrimSpace(matches[1])
	}
	return ""
}

package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

var apiBaseUrl = "http://pubs.vmware.com/vsphere-55/topic/com.vmware.wssdk.apiref.doc"

var mosIndex = apiBaseUrl + "/index-mo_types.html"
var doIndex = apiBaseUrl + "/index-do_types.html"

type Property struct {
	Name               string `json:"name"`
	Type_              string `json:"type"`
	Description        string `json:"description"`
	Optional           bool   `json:"optional"`
	RequiredPrivileges string `json:"required_privileges"`
}

type Parameter struct {
	Name               string `json:"name"`
	Type_              string `json:"type"`
	Description        string `json:"description"`
	Optional           bool   `json:"optional"`
	RequiredPrivileges string `json:"required_privileges"`
}

type Value struct {
	Type_       string `json:"type"`
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
}

var objDescRegex *regexp.Regexp

func init() {
	// if os.Getenv("GOMAXPROCS") == "" {
	// 	runtime.GOMAXPROCS(runtime.NumCPU())
	// }

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
	var wg sync.WaitGroup

	totalMos := 0
	totalDos := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		totalMos = scrape(mosIndex)
	}()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	totalDos = scrape(doIndex)
	// }()
	wg.Wait()

	log.Println("\n")
	log.Printf("Managed Objects: %d\n", totalMos)
	log.Printf("Data Objects: %d\n", totalDos)
	log.Println("Done 💩")
}

var closeChannel bool = false

func scrape(url string) int {
	d, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln(err)
	}

	totalObjects := 0
	channel := make(chan *Object)

	handler := func(i int, sel *goquery.Selection) {
		href, _ := sel.Attr("href")
		name, _ := sel.Attr("title")

		totalObjects++

		go scrapeObject(href, name, channel)
	}

	d.Find("#AE > nobr > a").Each(handler)
	//d.Find("#FJ > nobr > a").Each(handler)
	// d.Find("#KO > nobr > a").Each(handler)
	// d.Find("#PT > nobr > a").Each(handler)
	// d.Find("#UZ > nobr > a").Each(handler)

	for _ = range channel {
		if closeChannel {
			break
		}
		// log.Println("===============")
		// log.Printf("%v\n", obj)
		// log.Println("===============")
	}

	return totalObjects
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

	log.Println("type: " + name)
	//log.Println(obj.Description)

	d.Find(`body > table`).Each(func(i int, sel *goquery.Selection) {
		prev := strings.TrimSpace(sel.Prev().Text())
		if prev == "Properties" {
			obj.Properties = make([]*Property, sel.Find(`tbody > tr`).Length()-1)

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
						p.Type_ = strings.TrimSpace(type_.Last().Text())
					} else {
						p.Type_ = strings.TrimSpace(type_.Text())
					}

					p.Description = strings.TrimSpace(sel2.Find("td:nth-child(3)").Text())

					if p.Name != "" {
						obj.Properties = append(obj.Properties, p)
						log.Printf("%#v\n", p)
					}
				}
			})
		} else if prev == "Methods" {
			methods := sel.Find("tbody > tr:nth-child(2) > td > a")
			obj.Methods = make([]*Method, methods.Length())

			methods.Each(func(i int, sel2 *goquery.Selection) {
				mname := sel2.Text()
				m := &Method{}
				m.Name = mname

				log.Println("method: " + mname)

				m.Description = getMethodDesc(mname, html)

				d.Find("h1").Each(func(j int, sel3 *goquery.Selection) {
					if strings.TrimSpace(sel3.Text()) != mname {
						return
					}

					//Parameters
					sel4 := sel3.NextFilteredUntil(`p[class="table-title"]`, `table[cellspacing="0"]`).Next()

					m.Parameters = make([]*Parameter, sel4.Find(`tbody > tr`).Length()-1)

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

						p.Type_ = strings.TrimSpace(sel5.Find("td:nth-child(2) > a").Text())
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
							v.Type_ = strings.TrimSpace(type_.Last().Text())
						} else {
							v.Type_ = strings.TrimSpace(type_.Text())
						}

						v.Description = sel6.Find("td:nth-child(2)").Text()
						if v.Type_ != "" {
							m.ReturnValue = v
							//log.Printf("%#v\n", v)
						}
					})

					//Faults
					sel6 := sel5.NextFilteredUntil(`p[class="table-title"]`, `table[cellspacing="0"]`).Next()
					m.Faults = make([]Value, sel6.Find(`tbody > tr`).Length()-1)

					sel6.Find("tbody > tr").Each(func(k int, sel7 *goquery.Selection) {
						f := Value{}

						type_ := sel7.Find("td:nth-child(1) > a")
						if type_.Length() == 2 {
							f.Type_ = strings.TrimSpace(type_.Last().Text())
						} else {
							f.Type_ = strings.TrimSpace(type_.Text())
						}

						f.Description = sel7.Find("td:nth-child(2)").Text()
						if f.Type_ != "" {
							m.Faults = append(m.Faults, f)
							log.Printf("%#v\n", f)
						}
					})
				})
				obj.Methods = append(obj.Methods, m)
			})
		}
	})

	channel <- obj
	//closeChannel = true
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

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sync"
)

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

type Constant struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Object struct {
	Name        string      `json:"name"`
	Namespace   string      `json:"namespace"`
	Extends     string      `json:"extends"`
	Since       string      `json:"since"`
	Description string      `json:"description"`
	Properties  []*Property `json:"properties"`
	Methods     []*Method   `json:"methods"`
	Constants   []Constant  `json:"constants"`
}

var objDescRegex *regexp.Regexp

const (
	apiBase    = "./reference"
	moIndex    = apiBase + "/index-mo_types.html"
	doIndex    = apiBase + "/index-do_types.html"
	enumIndex  = apiBase + "/index-e_types.html"
	faultIndex = apiBase + "/index-faults.html"
)

func init() {
	/**
	* HTML is not well formed so we need this regexp to extract
	* Objects descriptions.
	 */
	objDescRegex = regexp.MustCompile(`</h2>((?s).+?)<p class="table-title">`)
}

type Counter struct {
	sync.Mutex
	x int
}

func scrape() {
	channel := make(chan *Object, 1)
	counter := &Counter{x: 0}
	totalObjects := 0
	var objs []*Object

	//ManagedObjects
	go func() {
		total := scrapeIndex(moIndex, "mo", channel)
		counter.Lock()
		counter.x += total
		counter.Unlock()
	}()

	//DataObjects
	go func() {
		total := scrapeIndex(doIndex, "do", channel)
		counter.Lock()
		counter.x += total
		counter.Unlock()
	}()

	//Enums
	go func() {
		total := scrapeIndex(enumIndex, "enum", channel)
		counter.Lock()
		counter.x += total
		counter.Unlock()
	}()

	//Faults
	go func() {
		total := scrapeIndex(faultIndex, "fault", channel)
		counter.Lock()
		counter.x += total
		counter.Unlock()
	}()

	for obj := range channel {
		totalObjects++

		log.Println("type: " + obj.Name)
		objs = append(objs, obj)

		counter.Lock()
		if totalObjects == counter.x {
			break
		}
		counter.Unlock()
	}

	//Sorts data api in order
	//to keep easily track of
	//changes through
	//a diff tool
	name := func(o1, o2 *Object) bool {
		return o1.Name < o2.Name
	}
	By(name).Sort(objs)

	data, err := json.MarshalIndent(objs, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	err = ioutil.WriteFile("./api.json", data, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("\nTotal objects to extract: %d\n", counter.x)
	log.Printf("Total objects extracted: %d\n", totalObjects)
	log.Println("vSphere API definition was generated in ./api.json")
	log.Println("Done 💩")
}

func scrapeIndex(refFile, namespace string, channel chan *Object) int {
	data, err := ioutil.ReadFile(refFile)
	if err != nil {
		log.Fatalln(err)
	}
	buffer := bytes.NewBuffer(data)

	d, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		log.Fatalln(err)
	}

	handler := func(i int, sel *goquery.Selection) {
		href, _ := sel.Attr("href")
		name, _ := sel.Attr("title")

		go scrapeObject(apiBase+"/"+href, name, namespace, channel)
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

func scrapeObject(refFile, name, namespace string, channel chan *Object) {
	data, err := ioutil.ReadFile(refFile)
	if err != nil {
		log.Fatalln(err)
	}
	buffer := bytes.NewBuffer(data)

	d, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		log.Fatalln(err)
	}

	obj := &Object{
		Name:      name,
		Namespace: namespace,
	}

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

	d.Find(`body > table`).Each(func(i int, sel *goquery.Selection) {
		prev := strings.TrimSpace(sel.Prev().Text())

		if prev == "Properties" {

			sel.Find("tbody > tr").Each(func(j int, sel2 *goquery.Selection) {
				if !sel2.HasClass("r0") && !sel2.HasClass("r1") {
					return
				}

				p := &Property{}
				pname := sel2.Find("td:nth-child(1)")

				pnameSel := pname.Find("strong")
				if pnameSel.Length() > 1 {
					p.Name = pnameSel.First().Text()
				} else {
					p.Name = pnameSel.Text()
				}

				p.Name = strings.TrimSpace(p.Name)

				if p.Name == "" {
					return
				}

				pname.Find("span").Each(func(k int, sel3 *goquery.Selection) {
					spanValue := strings.TrimSpace(sel3.Text())
					if spanValue == "*" {
						p.Optional = true
					} else if spanValue == "P" {
						p.RequiredPrivileges, _ = sel3.Attr("title")
					}
				})

				type_ := sel2.Find("td:nth-child(2) > a")
				if type_.Length() == 1 {
					p.Type = strings.TrimSpace(type_.Text())
				} else if type_.Length() == 2 {
					//Instead of getting ManagedReferenceObject, it gets
					//the real type
					p.Type = strings.TrimSpace(type_.Last().Text())
				} else {
					//Gets primitive types such as: xsd:string, xsd:long, etc
					typeSel := sel2.Find("td:nth-child(2)")

					if typeSel.Length() > 1 {
						p.Type = typeSel.First().Text()
					} else {
						p.Type = typeSel.Text()
					}

					p.Type = strings.TrimSpace(p.Type)
				}

				p.Description = strings.TrimSpace(sel2.Find("td:nth-child(3)").Text())

				obj.Properties = append(obj.Properties, p)

			})

			//In order to respect WSDL sequences, we have to
			//read the WSDL type definition for every data type,
			//and make sure the order of the properties is the same
			//in the resulting api.json definition.
			wsdlDef := d.Find("#wsdl-div > textarea").Text()

			wsdl := &XsdComplexType{}
			xml.Unmarshal([]byte(wsdlDef), wsdl)

			props := make([]*Property, 0)
			for _, el := range wsdl.ComplexContent.Extension.Sequence.Elements {
				for _, prop := range obj.Properties {
					if prop.Name == el.Name {
						props = append(props, prop)
						break
					}
				}
			}

			if len(props) > 0 {
				obj.Properties = props
			}
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
						if !sel5.HasClass("r0") && !sel5.HasClass("r1") {
							return
						}
						p := &Parameter{}
						pname := sel5.Find("td:nth-child(1)")

						pnameSel := pname.Find("strong")
						if pnameSel.Length() > 1 {
							p.Name = pnameSel.First().Text()
						} else {
							p.Name = pnameSel.Text()
						}

						p.Name = strings.TrimSpace(p.Name)

						pname.Find("span").Each(func(k int, sel6 *goquery.Selection) {
							spanValue := strings.TrimSpace(sel6.Text())
							if spanValue == "*" {
								p.Optional = true
							} else if spanValue == "P" {
								p.RequiredPrivileges, _ = sel6.Attr("title")
							}
						})

						type_ := sel5.Find("td:nth-child(2) > a")
						if type_.Length() == 1 {
							p.Type = strings.TrimSpace(type_.Text())
						} else if type_.Length() == 2 {
							//Instead of getting ManagedReferenceObject, it gets
							//the real type
							p.Type = "mo." + strings.TrimSpace(type_.Last().Text())
						} else {
							//Gets primitive types such as: xsd:string, xsd:long, etc
							typeSel := sel5.Find("td:nth-child(2)")

							if typeSel.Length() > 1 {
								p.Type = typeSel.First().Text()
							} else {
								p.Type = typeSel.Text()
							}

							p.Type = strings.TrimSpace(p.Type)

						}

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
						if v.Type != "" && v.Type != "None" {
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

//This is required to conform with godoc
//convention
func sanitizeDesc(text string) string {
	if text == "" {
		return text
	}

	s := strings.Replace(text, "<br>", "\n", -1)
	s = strings.Replace(s, "</br>", "\n", -1)
	s = strings.Replace(s, "<p></p>", "", -1)
	s = strings.Replace(s, "<p>", "", -1)
	s = strings.Replace(s, "</p>", "\n", -1)
	s = strings.Replace(s, "<ul>", "\n", -1)
	s = strings.Replace(s, "</ul>", "\n", -1)
	s = strings.Replace(s, "<li>", "• ", -1)
	s = strings.Replace(s, "</li>", "\n", -1)
	s = strings.Replace(s, `<a id="field_detail" name="field_detail"></a>`, "", -1)

	// Remove a few common harmless entities, to arrive at something more like plain text
	// This relies on having removed *all* tags above
	s = strings.Replace(s, "&nbsp;", " ", -1)
	s = strings.Replace(s, "&quot;", "\"", -1)
	s = strings.Replace(s, "&apos;", "'", -1)
	s = strings.Replace(s, "&#34;", "\"", -1)
	s = strings.Replace(s, "&#39;", "'", -1)
	// NB spaces here are significant - we only allow & not part of entity
	s = strings.Replace(s, "&amp; ", "& ", -1)
	s = strings.Replace(s, "&amp;amp; ", "& ", -1)

	b := bytes.NewBufferString("")

	//Remove remaining tags
	inTag := false
	for _, r := range s {
		switch r {
		case '<':
			inTag = true
		case '>':
			inTag = false
		default:
			if !inTag {
				b.WriteRune(r)
			}
		}
	}
	s = b.String()

	return s
}

func getObjectDesc(html string) string {
	matches := objDescRegex.FindStringSubmatch(html)
	if len(matches) == 2 {
		return sanitizeDesc(strings.TrimSpace(matches[1]))
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
		return sanitizeDesc(strings.TrimSpace(matches[1]))
	}
	return ""
}

package vim

import (
	"encoding/xml"
)

//Use to unmarshal vSphere dynamic types
type PropertyValue struct {
	Type string `xml:"type,attr"`
	Xml  string `xml:",innerxml"`
}

func (pv *PropertyValue) Value() (interface{}, error) {
	v := registry[pv.Type]()
	err := xml.Unmarshal([]byte("<val>"+pv.Xml+"</val>"), v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

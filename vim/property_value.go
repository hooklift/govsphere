// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
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

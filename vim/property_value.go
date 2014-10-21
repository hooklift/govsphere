// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vim

import (
	"encoding/xml"
	"fmt"
)

//Use to unmarshal vSphere dynamic types
type PropertyValue struct {
	Type string `xml:"type,attr"`
	Xml  string `xml:",innerxml"`
}

func (pv *PropertyValue) Value() (interface{}, error) {
	vimType := pv.Type

	if getTypeInstance, ok := registry[vimType]; ok {
		v := getTypeInstance()

		data := []byte("<val>" + pv.Xml + "</val>")
		err := xml.Unmarshal(data, v)
		if err != nil {
			return nil, err
		}

		return v, nil
	}

	return nil, fmt.Errorf("vSphere type not found in registry: %s", vimType)
}

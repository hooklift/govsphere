// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vim

import (
	"errors"
	"log"
)

type ManagedObject struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",innerxml"`
}

func (mo *ManagedObject) currentProperty(property string) (interface{}, error) {
	if session == nil {
		return nil, errors.New("A vSphere session must be created first")
	}

	this := &ManagedObjectReference{
		Type:  mo.Type,
		Value: mo.Value,
	}

	//Create object spec
	objSpec := &ObjectSpec{
		Obj:  this,
		Skip: false,
	}

	//If an empty property is received
	//it reads all properties from the object
	//and PathSet is ignored.
	allProps := false
	if property == "" {
		allProps = true
	}

	//Create property spec
	propSpec := &PropertySpec{
		Type:    mo.Type,
		All:     allProps,
		PathSet: []string{property},
	}

	//Create property filter spec using object and property specs
	filterSpec := &PropertyFilterSpec{
		PropSet:   []*PropertySpec{propSpec},
		ObjectSet: []*ObjectSpec{objSpec},
	}

	//Retrieves ServiceContent using method instead of Content() accessor
	//to avoid a stack overflow
	pc := session.ServiceContent.PropertyCollector

	//It does not use pc.RetrievePropertiesEx because we want
	//to support older versions of vSphere too
	objs, err := pc.RetrieveProperties([]*PropertyFilterSpec{filterSpec})
	if err != nil {
		return nil, err
	}

	if len(objs) == 0 ||
		len(objs[0].MissingSet) > 0 ||
		len(objs[0].PropSet) == 0 {
		return nil, nil
	}

	log.Printf("Returned val: %#v\n", objs[0].PropSet[0].Val)
	return objs[0].PropSet[0].Val.Value()
}

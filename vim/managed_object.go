// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vim

import (
	"errors"
)

type ManagedObject struct {
	*ManagedObjectReference
}

func (mo *ManagedObject) currentProperty(name string) (interface{}, error) {
	if session == nil {
		return nil, errors.New("A vSphere session must be created first")
	}

	if name == "" {
		return nil, errors.New("A property name is required")
	}

	this := &ManagedObjectReference{
		Type:  mo.Type,
		Value: mo.Value,
	}

	// Create object spec
	objSpec := &ObjectSpec{
		Obj:  this,
		Skip: false,
	}

	// Create property spec
	propSpec := &PropertySpec{
		Type:    mo.Type,
		All:     false,
		PathSet: []string{name},
	}

	// Create property filter spec using object and property specs
	filterSpec := &PropertyFilterSpec{
		PropSet:   []*PropertySpec{propSpec},
		ObjectSet: []*ObjectSpec{objSpec},
	}

	// Retrieves ServiceContent using method instead of Content() accessor
	// to avoid a stack overflow
	pc := session.ServiceContent.PropertyCollector

	// It does not use pc.RetrievePropertiesEx because we want to support older
	// versions of vSphere too
	objs, err := pc.RetrieveProperties([]*PropertyFilterSpec{filterSpec})
	if err != nil {
		return nil, err
	}

	if len(objs) == 0 ||
		len(objs[0].MissingSet) > 0 ||
		len(objs[0].PropSet) == 0 {
		return nil, nil
	}

	val, err := objs[0].PropSet[0].Val.Value()
	return val, err
}

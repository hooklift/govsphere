package vim

import (
	"errors"
	"log"
)

type ManagedObject struct {
	Type           string `xml:"type,attr"`
	Value          string `xml:",innerxml"`
	serviceIntance *ServiceInstance
}

func (mo *ManagedObject) currentProperty(property string, response interface{}) error {
	if serviceInstance == nil {
		return errors.New("A ServiceInstance must be created first")
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
	serviceContent, err := serviceInstance.RetrieveServiceContent()
	if err != nil {
		return err
	}

	pc := serviceContent.PropertyCollector
	log.Println("================================================")
	log.Printf("%#v\n", pc)
	log.Println("================================================")

	//It does not use RetrievePropertiesEx because we want
	//to support older versions of vSphere too, when getting
	//MO properties
	objs, err := pc.RetrieveProperties([]*PropertyFilterSpec{filterSpec})
	if err != nil {
		return err
	}

	log.Printf("%#v\n", objs)

	//Get property set from returned ObjectContent
	return nil
}

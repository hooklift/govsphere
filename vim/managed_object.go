package vim

import (
	"github.com/c4milo/govsphere/vim/soap"
)

type ManagedObject struct {
	this       *ManagedObjectReference `xml:"_this"`
	soapClient *soap.Client
}

func (mo *ManagedObject) currentProperty(property string, response interface{}) error {
	return nil
}

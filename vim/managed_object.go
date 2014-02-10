package vim

import (
	"github.com/c4milo/govsphere/vim/soap"
)

type ManagedObject struct {
	This       *ManagedObjectReference
	soapClient *soap.Client
}

func (mo *ManagedObject) currentProperty(property string, response interface{}) error {
	return nil
}

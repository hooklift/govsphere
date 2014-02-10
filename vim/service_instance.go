package vim

import (
	"github.com/c4milo/govsphere/vim/soap"
)

const (
	APIv4_0 string = "urn:vim25/4.0"
	APIv4_1 string = "urn:vim25/4.1"
	APIv5_0 string = "urn:vim25/5.0"
	APIv5_1 string = "urn:vim25/5.1"
	APIv5_5 string = "urn:vim25/5.5"
)

func NewServiceInstance(url, user, pass, apiVersion string, ignoreCert bool) *ServiceInstance {
	if url == "" || user == "" {
		panic("A URL and username is required")
	}

	service := &ServiceInstance{
		ManagedObject: &ManagedObject{
			this: &ManagedObjectReference{
				Type:  "ServiceInstance",
				Value: "ServiceInstance",
			},
			soapClient: soap.NewClient(url, apiVersion, ignoreCert),
		},
	}

	return service
}

package vim

import (
	"github.com/c4milo/govsphere/vim/soap"
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

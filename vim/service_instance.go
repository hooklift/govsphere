package vim

import (
	"github.com/c4milo/govsphere/vim/soap"
)

var soapClient *soap.Client
var serviceInstance *ServiceInstance

const (
	APIv4_0 string = "urn:vim25/4.0"
	APIv4_1 string = "urn:vim25/4.1"
	APIv5_0 string = "urn:vim25/5.0"
	APIv5_1 string = "urn:vim25/5.1"
	APIv5_5 string = "urn:vim25/5.5"
)

func NewServiceInstance(url, user, pass string, ignoreCert bool) *ServiceInstance {
	if serviceInstance != nil {
		return serviceInstance
	}

	if url == "" {
		panic("Server URL is required")
	}

	service := &ServiceInstance{
		ManagedObject: &ManagedObject{
			Type:  "ServiceInstance",
			Value: "ServiceInstance",
		},
	}

	//Since we do not know the latest version supported by
	//the vSphere server yet, we use the oldest possible first.
	soapClient = soap.NewClient(url, APIv4_0, ignoreCert)

	sc, err := service.RetrieveServiceContent()
	if err != nil {
		panic(err)
	}

	version := sc.About.ApiVersion

	var apiVersion string
	switch version {
	default:
		apiVersion = APIv5_5
	case "4.0":
		apiVersion = APIv4_0
	case "4.1":
		apiVersion = APIv4_1
	case "5.0":
		apiVersion = APIv5_0
	case "5.1":
		apiVersion = APIv5_1
	case "5.5":
		apiVersion = APIv5_5
	}

	//Now that we know the latest supported API version,
	//we can re-create soapClient using such version.
	soapClient = soap.NewClient(url, apiVersion, ignoreCert)

	serviceInstance = service

	sc, err = service.RetrieveServiceContent()
	if err != nil {
		panic(err)
	}

	// session, err := sc.SessionManager.Login(user, pass, "")
	// if err != nil {
	// 	panic(err)
	// }

	//log.Printf("%v\n", session)

	return service
}

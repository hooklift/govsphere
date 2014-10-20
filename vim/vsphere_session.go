// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vim

import (
	"net/http"

	"github.com/cloudescape/govsphere/vim/soap"
)

type VSphereSession struct {
	sessionId       string
	soapClient      *soap.Client
	ServiceInstance *ServiceInstance
	ServiceContent  *ServiceContent
	UserSession     *UserSession
}

// For now, let's only support
// one vSphere session per App
var session *VSphereSession

const (
	APIv4_0 string = "urn:vim25/4.0"
	APIv4_1 string = "urn:vim25/4.1"
	APIv5_0 string = "urn:vim25/5.0"
	APIv5_1 string = "urn:vim25/5.1"
	APIv5_5 string = "urn:vim25/5.5"
)

func NewVSphereSession(url, user, pass string, ignoreCert bool) *VSphereSession {
	// For now, let's only support
	// one vSphere session per App
	if session != nil {
		return session
	}

	if url == "" {
		panic("Server URL is required")
	}

	objref := &ManagedObjectReference{
		Type:  "ServiceInstance",
		Value: "ServiceInstance",
	}

	service := &ServiceInstance{
		ManagedObject: &ManagedObject{
			ManagedObjectReference: objref,
		},
	}

	// Since we do not know the latest version supported by
	// the vSphere server yet, we use the oldest possible first.
	soapClient := soap.NewClient(url, APIv4_0, ignoreCert)

	session = &VSphereSession{
		soapClient: soapClient,
	}

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

	// Now that we know the latest supported API version,
	// we can re-create soapClient using such version.
	session.soapClient = soap.NewClient(url, apiVersion, ignoreCert)

	sc, err = service.RetrieveServiceContent()
	if err != nil {
		panic(err)
	}

	userSession, err := sc.SessionManager.Login(user, pass, "")
	if err != nil {
		panic(err)
	}

	session.ServiceInstance = service
	session.ServiceContent = sc
	session.UserSession = userSession

	return session
}

func (s *VSphereSession) invoke(request interface{}, response interface{}) error {
	// Sets session cookie
	var cookie *http.Cookie
	if s.sessionId != "" {
		cookie = &http.Cookie{
			Name:     "vmware_soap_session",
			Value:    s.sessionId,
			Secure:   true,
			HttpOnly: true,
			Path:     "/",
		}
	}

	cookies, err := s.soapClient.Call(request, response, []*http.Cookie{cookie})
	if err != nil {
		return err
	}

	if len(cookies) > 0 {
		s.sessionId = cookies[0].Value
	}
	return nil
}

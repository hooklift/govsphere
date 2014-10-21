// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package main

import (
	"fmt"

	"github.com/cloudescape/govsphere/vim"
)

func main() {
	session := vim.NewVSphereSession("https://172.16.103.128/sdk", "user", "password", true)

	c, err := session.ServiceInstance.Capability()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Capabilities -> %+v\n", c)

	// Supported since vSphere API 5.1 only
	// userAgent := session.UserSession.UserAgent
	// fmt.Printf("User Agent -> %s\n", userAgent)

	err = session.ServiceContent.SessionManager.Logout()
	if err != nil {
		panic(err)
	}

	c, err = session.ServiceInstance.Capability()
	if err != nil {
		panic(err)
	}

	fmt.Printf("There is not an active session, Capabilities -> %+v\n", c)
}

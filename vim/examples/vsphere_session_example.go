// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package examples

import (
	"fmt"
	"testing"

	"github.com/cloudescape/govsphere/vim"
)

// TODO(c4milo): Remove after Go 1.4.
// Related to https://codereview.appspot.com/107320046
func TestB(t *testing.T) {}

func Example_2() {
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

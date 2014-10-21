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
func TestA(t *testing.T) {}

func Example_1() {

	session := vim.NewVSphereSession("https://172.16.103.128/sdk", "username", "password", true)

	content, err := session.ServiceInstance.Content()

	if err != nil {
		panic("unable to get the service content")
	}

	children, err := content.RootFolder.ChildEntity() // <------- THIS CALL WILL PANIC

	if err != nil {
		panic(err)
	}

	fmt.Printf("--> %#v <--\n", children[0].Type)

	types, err := content.RootFolder.ChildType()
	if err != nil {
		panic(err)
	}
	fmt.Printf("--> %#v <--\n", types)

	err = session.ServiceContent.SessionManager.Logout()
	if err != nil {
		panic("unable to logout from the vsphere server")
	}

	fmt.Println("Done!")
}

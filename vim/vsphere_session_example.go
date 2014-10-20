// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package vim

import "testing"

func TestNewVSphereSession(t *testing.T) {
	session := NewVSphereSession("https://<VCenter or ESXi IP address>/sdk", "user", "password", true)

	_, err := session.ServiceInstance.Capability()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	// Supported since vSphere API 5.1 only
	// userAgent := session.UserSession.UserAgent
	// if !strings.Contains(userAgent, "govsphere") {
	// 	t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", userAgent, "govsphere")
	// }

	err = session.ServiceContent.SessionManager.Logout()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	c, err := session.ServiceInstance.Capability()
	if c != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", c, nil)
	}
}

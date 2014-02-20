package vim

import (
	"strings"
	"testing"
)

func TestNewVSphereSession(t *testing.T) {
	session := NewVSphereSession("https://172.16.103.128/sdk", "root", "shadow04101", true)

	_, err := session.ServiceInstance.Capability()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	userAgent := session.UserSession.UserAgent
	if !strings.Contains(userAgent, "govsphere") {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", userAgent, "govsphere")
	}

	err = session.ServiceContent.SessionManager.Logout()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	c, err := session.ServiceInstance.Capability()
	if c != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", c, nil)
	}
}

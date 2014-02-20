package vim

import (
	"testing"
)

func TestNewVSphereSession(t *testing.T) {
	session := NewVSphereSession("https://172.16.103.128/sdk", "root", "shadow04101", true)

	cap, err := session.ServiceInstance.Capability()
	if err != nil {
		t.Errorf("incorrect result\ngot:  %#v\nwant: %#v", err, nil)
	}

	t.Logf("=>%#v<=\n", cap)
}

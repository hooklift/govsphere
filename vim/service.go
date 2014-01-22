package vim

import (
	"sync"
)

type serviceInstance struct {
	url, user, pass string
	tls             bool
}

var once sync.Once
var instance *serviceInstance

func NewServiceInstance(url, user, pass string, tls bool) *serviceInstance {
	//Only one instance of ServiceInstance is allowed.
	once.Do(func() {
		instance = &serviceInstance{
			url:  url,
			user: user,
			pass: pass,
			tls:  tls,
		}
	})

	return instance
}

/*
Fully documented and friendly VMWare vSphere API for Go.

Versions supported

vSphere 5.5 and previous versions.

Prerequisite

Before attempting to use this API please take some time to familiarize yourself
with [VMware VI object model](http://www.doublecloud.org/2010/02/object-model-of-vmware-vsphere-api-a-big-picture-in-2-minutes/)

CLI installation

 $ go install

API code generation

 $ govsphere generate

API definitions generation

The generation process is going to create a file called api.json, relative to the path
from where the command is executed.

 $ govsphere scrape

API Documentation

http://godoc.org/github.com/cloudescape/govsphere
*/
package vim

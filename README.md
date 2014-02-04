# Go API for VMware vSphere

Fully documented and friendly VMWare vSphere API for Go.

### Versions supported
5.5 and previous versions 

## Prerequisite
Before attempting to use this API please take some time to familiarize yourself with [VMware VI object model](http://www.doublecloud.org/2010/02/object-model-of-vmware-vsphere-api-a-big-picture-in-2-minutes/) 


## CLI installation
`go install`

## API code generation
`$ govsphere generate`

## API definitions generation
Please be aware that this procedure makes several HTTP requests to http://pubs.vmware.com. 
The generation process is going to create a file called api.json, relative to the path 
from where the command was executed.

`$ govsphere scrape`

Example output: 

![](https://i.cloudup.com/sEiUVA9a8L.png)

## Documentation
TODO

## TODO
* Create ServiceInstance class
	* Use best API version supported by the server
* LoginByToken
* Add support for JSON Schema in api.json generation

## License
Copyright 2014 Cloudescape. All rights reserved.

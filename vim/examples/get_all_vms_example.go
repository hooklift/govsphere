package main

import (
	"fmt"
	"github.com/cloudescape/govsphere/vim"
	"io/ioutil"
	"log"
)

type EntityProcessor func(*vim.ManagedEntity)

func FolderTraverser(folder *vim.Folder, targetEntity string, processor EntityProcessor) {
	children, err := folder.ChildEntity()

	if err != nil {
		fmt.Printf("--> Error occurred while getting children entities for folder: %#v\n", err)
	} else {

		for _, child := range children {

			if child.Type == targetEntity {
				processor(child)
			} else if child.Type == "Folder" {

				folder := &vim.Folder{
					ManagedEntity: child,
				}

				FolderTraverser(folder, targetEntity, processor)

			} else {
				continue
			}

		}

	}

	return
}

func ProcessVm(entity *vim.ManagedEntity) {

	vm := &vim.VirtualMachine{
		ManagedEntity: entity,
	}

	summary, err := vm.Summary()

	if err != nil {
		fmt.Printf("--> Encountered an error while getting VM summary: %#v <--\n", err)
		return
	}

	fmt.Printf("Name       : %s\n", summary.Config.Name)
	fmt.Printf("Path       : %s\n", summary.Config.VmPathName)
	fmt.Printf("Guest      : %s\n", summary.Config.GuestFullName)
	fmt.Printf("Annotation : %s\n", summary.Config.Annotation)
	fmt.Printf("State      : %s\n", *summary.Runtime.PowerState)

	if summary.Guest != nil {
		ip := summary.Guest.IpAddress
		fmt.Printf("IP         : %s\n", ip)
	}

	if summary.Runtime.Question != nil {
		fmt.Printf("Question   : %s\n", summary.Runtime.Question.Text)
	}

	fmt.Printf("\n")
}

func main() {

	//disable logging.
	log.SetOutput(ioutil.Discard)

	session := vim.NewVSphereSession("https://myhost.com/sdk", "myusername", "mypassword", true)

	content, err := session.ServiceInstance.Content()

	if err != nil {
		fmt.Printf("--> Encountered an error getting content from ServiceInstance: %#v <--\n", err)
		panic(err)
	}

	children, err := content.RootFolder.ChildEntity()

	if err != nil {
		fmt.Printf("--> Encountered an error getting childEntity from vmFolder: %#v <--\n", err)
		panic(err)
	}

	for _, child := range children {

		if child.Type == "Datacenter" {

			datacenter := &vim.Datacenter{
				ManagedEntity: child,
			}

			vmfolder, err := datacenter.VmFolder()
			if err != nil {
				fmt.Printf("--> Encountered an error getting vmFolder from datacenter: %#v <--\n", err)
				panic(err)
			}

			FolderTraverser(vmfolder, "VirtualMachine", ProcessVm)
		}
	}

	err = session.ServiceContent.SessionManager.Logout()

	if err != nil {
		fmt.Printf("--> Unable to logout! Error: %#v <--\n", err)
		panic(err)
	}

	fmt.Println("Done!")
}

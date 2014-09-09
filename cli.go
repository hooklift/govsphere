// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package main

import (
	"log"
	"os"
	"runtime"
)

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func init() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	log.SetFlags(0)
	log.SetOutput(os.Stdout)
}

func usage() {
	log.Println(`
Usage:
  ./govsphere [COMMAND]

Application Commands:
• generate: Generates Go code from API definitions
• scrape  : Generates API definitions into ./api.json file
`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	action := os.Args[1]

	if action == "scrape" {
		scrape()
	} else if action == "generate" {
		generate("./api.json")
	} else {
		usage()
		os.Exit(1)
	}
}

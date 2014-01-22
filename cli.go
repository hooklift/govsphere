package main

import (
	"log"
	"os"
	"runtime"
)

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

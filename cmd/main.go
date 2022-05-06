package main

import (
	"flag"
	"fmt"
	"os"
	"vector-configurator/pkg/configurator"
)

func main() {

	// set environment arguments
	setenv := flag.NewFlagSet("set-environment", flag.ExitOnError)
	shost := setenv.String("h", "", "The internal IP address of the Vector robot")
	skey := setenv.String("k", "", "The location of the SSH key")
	senv := setenv.String("e", "", "environment")

	cloudbin := flag.NewFlagSet("upload-cloud-binaries", flag.ExitOnError)
	chost := cloudbin.String("h", "", "The internal IP address of the Vector robot")
	ckey := cloudbin.String("k", "", "The location of the SSH key")
	cbindir := cloudbin.String("b", "", "Location of / path to vic-cloud and vic-gateway binaries")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println(`vector configurator

This tool will allow you to do the following things:
set-environment          - change the environment that 
                           your bot is pointed at
upload-cloud-binaries    - upload and set permissions
                           for the vector-cloud binaries`)
		os.Exit(0)
	}

	switch os.Args[1] {
	case "set-environment":
		_ = setenv.Parse(os.Args[2:])
		if *senv == "" || *shost == "" || *skey == "" {
			setenv.Usage()
			os.Exit(1)
		}
		configurator.SetEnvironment(*shost, *skey, *senv)
	case "upload-cloud-binaries":
		_ = cloudbin.Parse(os.Args[2:])
		if *cbindir == "" || *chost == "" || *ckey == "" {
			cloudbin.Usage()
			os.Exit(1)
		}
		configurator.UploadCloud(*chost, *ckey, *cbindir)

	}

}

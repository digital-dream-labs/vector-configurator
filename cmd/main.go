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
	shost := setenv.String("h", "", "the ip of the vector robot")
	skey := setenv.String("k", "", "the location of the ssh key")
	senv := setenv.String("e", "", "environment")

	cloudbin := flag.NewFlagSet("upload-cloud-binaries", flag.ExitOnError)
	chost := cloudbin.String("h", "", "the ip of the vector robot")
	ckey := cloudbin.String("k", "", "the location of the ssh key")
	cbindir := cloudbin.String("b", "", "path to vic-cloud and vic-gateway binaries")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println(`vector configurator

This tool will allow you to do the following things:
set-environment          - change the environment that 
                           your bot is pointed at
upload-cloud-binaries    - upload and set permissions
                           for the vector-cloud binaries
`)
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

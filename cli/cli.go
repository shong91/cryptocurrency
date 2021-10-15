package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/shong91/cryptocurrency/explorer"
	"github.com/shong91/cryptocurrency/rest"
)

func usage() {
	fmt.Printf("Welcome to cryptocurrency \n\n")
	fmt.Printf("please use the following flags: \n\n")
	fmt.Printf("-port=4000: Set the PORT of the server \n")
	fmt.Printf("-mode=rest: Choose between 'html' and 'rest' \n")

	os.Exit(0) // error code 0 = none
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		// start REST api
		rest.Start(*port)
	case "html":
		// start html explorer
		explorer.Start(*port)
	default:
		usage()

	}

	fmt.Println(*port, *mode)
}
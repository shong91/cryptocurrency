package main

import (
	"flag"
	"fmt"
	"os"
)

func usage(){
	fmt.Printf("Welcome to cryptocurrency \n\n")
	fmt.Printf("please use the following command: \n\n")
	fmt.Printf("explorer: this starts the HTML Explorer \n")
	fmt.Printf("rest: this starts the REST API (recommended) \n")

	os.Exit(0) // error code 0 = none
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := rest.Int("port", 4000, "Sets the port of ther server")

	switch os.Args[1] {
	case "explorer":
		fmt.Printf("Start explorer")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println("Start server with port:", *portFlag)	
	}
}

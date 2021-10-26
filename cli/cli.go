package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/shong91/cryptocurrency/explorer"
	"github.com/shong91/cryptocurrency/rest"
)

func usage() {
	fmt.Printf("Welcome to cryptocurrency \n\n")
	fmt.Printf("please use the following flags: \n\n")
	fmt.Printf("-port=4000: Set the PORT of the server \n")
	fmt.Printf("-mode=rest: Choose between 'html' and 'rest' \n")

	// 모든 함수를 제거하지만, 그 전에 defer 를 먼저 이행한다.
	runtime.Goexit()

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
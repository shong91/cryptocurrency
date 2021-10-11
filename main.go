package main

import (
	"github.com/shong91/cryptocurrency/rest"
)


func main() {
	// go routine 을 사용하여 동시에 실행되도록 한다
	// go explorer.Start(3000)
	rest.Start(4000)

}
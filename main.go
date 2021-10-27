package main

import (
	"github.com/shong91/cryptocurrency/cli"
	"github.com/shong91/cryptocurrency/db"
)


func main() {
	defer db.Close()
	cli.Start()
}

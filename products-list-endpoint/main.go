package main

import (
	"github.com/zeroberto/products-store/products-list-endpoint/cmd"
)

func main() {

	server := new(cmd.Server)

	defer server.Stop()

	server.Start()
}

package main

import (
	"github.com/zeroberto/products-store/products-list-endpoint/cmd"
	"github.com/zeroberto/products-store/products-list-endpoint/container"
)

var c container.Container

func main() {

	server := new(cmd.Server)

	defer server.Stop()

	server.Start()
}

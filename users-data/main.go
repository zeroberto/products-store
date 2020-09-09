package main

import (
	"github.com/zeroberto/products-store/users-data/cmd"
)

func main() {
	server := new(cmd.Server)
	server.Start()
}

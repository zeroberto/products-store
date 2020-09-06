package main

import (
	"github.com/zeroberto/products-store/discount-calculator/cmd"
)

func main() {

	server := new(cmd.Server)

	defer server.Stop()

	server.Start()
}

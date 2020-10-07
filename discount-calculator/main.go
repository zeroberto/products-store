package main

import (
	"github.com/zeroberto/products-store/discount-calculator/cmd"
)

func main() {
	cmd.NewServer().Start()
}

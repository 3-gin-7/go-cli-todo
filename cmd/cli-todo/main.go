package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Starting todo cli")

	var n = flag.Int("n", 123, "help message for flag n")

	fmt.Printf("flag value: %v\r\n", *n)
	flag.Parse()
	fmt.Printf("flag value: %v\r\n", *n)
}

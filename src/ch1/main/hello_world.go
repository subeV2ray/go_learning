package main

import (
	"fmt"
	"os"
)

func main() {
	println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
}

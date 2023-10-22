package main

import (
	"fmt"
	"os"

	"github.com/nnamm/gomd/generator"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("expected a single argument")
		return
	}

	arg := os.Args[1]
	if err := generator.GenerateContent(arg); err != nil {
		fmt.Println(err)
		return
	}
}

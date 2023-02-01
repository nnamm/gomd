package main

import (
	"fmt"
	"os"
)

func main() {
	//if err := run(); err != nil {
	//	log.Printf("Failed to generate file: %v", err)
	//}
	if len(os.Args) < 2 {
		fmt.Println("Required args are missing.")
		return
	}
	dirName := os.Args[1]

	err := os.Mkdir(dirName, 0755)
	if os.IsExist(err) {
		fmt.Println("The directory already exists.")
		return
	}
	if err != nil {
		fmt.Printf("Fail to mkdir: %v", err)
		return
	}

	err = os.WriteFile(dirName+"/aaa.md", []byte("Hello, Gophers!"), 0644)
	if err != nil {
		fmt.Printf("Fail to mkdir: %v", err)
		return
	}
}

//var articleNum string
//
//func run() error {
//	flag.StringVar(&articleNum, "n", "000", "Article number of the markdown file to be newly created")
//
//	if err := validateArgs(); err != nil {
//		fmt.Fprintln(os.Stderr, err)
//		os.Exit(1)
//	}
//
//	flag.Parse()
//	fmt.Printf("param -n : %s\n", articleNum)
//
//	return nil
//}
//
//func validateArgs() error {
//	return nil
//}

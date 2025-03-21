package main

import (
	"fmt"
	"os"
)

func main() {
	argu := os.Args
	switch {
	case len(argu) < 2:
		fmt.Println("File name missing")
	case len(argu) > 2:
		fmt.Println("Too many arguments")
	default:
		filename := argu[1]
		contents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
		}
		os.Stdout.Write(contents)
	}
}

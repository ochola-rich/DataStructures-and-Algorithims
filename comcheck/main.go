package main

import (
	"fmt"
	"os"
)

func main() {
	argu := os.Args[1:]
	for _, substr := range argu {
		if substr == "01" || substr == "galaxy" || substr == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}

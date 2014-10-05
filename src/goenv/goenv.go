package main

import (
	"fmt"
	"os"
)

func main() {
	switch {
	case len(os.Args) >= 2:
		CreateGoEnv(os.Args[1])
	default:
		fmt.Println("You must provide a destination folder")
		Usage()
	}
}

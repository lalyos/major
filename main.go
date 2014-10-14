package main

import (
	"fmt"
	"os"
)

func main() {

	major := os.Getenv("USER")
	fmt.Println("The major of Budapest is: ...", major, " he rules \n")
}

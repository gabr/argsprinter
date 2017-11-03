package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Given args:                %v\n", os.Args)
	fmt.Printf("Args structure:            %#v\n", os.Args)
	fmt.Println("Args without program path:", os.Args[1:])
}

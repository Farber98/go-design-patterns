package main

import (
	"fmt"
	"os"

	"github.com/Farber98/go-design-patterns/structuralPatterns/decorator"
)

func main() {
	fmt.Println("Enter the type number of server you want to launch from the following:")
	fmt.Println("1.- Plain server")
	fmt.Println("2.- Server with logging")
	fmt.Println("3.- Server with logging and auth")
	var selection int
	fmt.Fscanf(os.Stdin, "%d", &selection)
	decorator.MyServer(selection)
}

// main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	shs := allShips()
	for val, i := range *shs {
		fmt.Println(val, i) //test of all things shipy. Will change soon
	}
}

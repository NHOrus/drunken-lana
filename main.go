// main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	allships := importShips()
	for _, val := range allships {
		fmt.Println(val)
	}
}

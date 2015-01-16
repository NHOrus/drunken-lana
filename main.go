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

	rule, _ := getRules()
	leftField := newField(rule.fLen, rule.fWid)
	rightField := newField(rule.fLen, rule.fWid)
	rfc := make(chan Field)
	rfc <- *rightField
	lfc := make(chan Field)
	lfc <- *leftField

	drawfield(lfc, rfc)
}

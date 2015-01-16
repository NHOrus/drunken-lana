// main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	rule, _ := getRules()

	shs := rule.Ships
	for val, i := range *shs {
		fmt.Println(val, i) //test of all things shipy. Will change soon
	}

	leftField := newField(rule.fLen, rule.fWid)
	rightField := newField(rule.fLen, rule.fWid)

	fmt.Println("Created fields")

	rfc := make(chan Field, 1)
	rfc <- rightField
	lfc := make(chan Field, 1)
	lfc <- leftField
	fmt.Println("Beginning to draw")
	drawfield(lfc, rfc)
}

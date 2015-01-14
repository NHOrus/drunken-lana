// field.go'
package main

import (
	"fmt"
)

const (
	empty = iota
	full
	miss
	hit
)

type Cell struct {
	state  int
	weight float64
}

type Field struct {
	Grid          *[][]Cell
	Length, Width int
}

func newField(length int, width int) (retField *Field) {
	tCell := make([][]Cell, length*width)
	retField.Grid = &tCell
	retField.Length = length
	retField.Width = width
	return
}

func (Fi *Field) putShips(shs map[string]ship, putter func(field *Field, fat int) error) error {
	for name, params := range shs {
		fmt.Println(name)
		for i := 0; i < params.Numbers; i++ {
			if err := putter(Fi, params.Fatness); err != nil {
				return err
			}
		}
	}

	return nil
}

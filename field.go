// field.go'
package main

const (
	empty = iota
	ship
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

func NewField(length int, width int) (retField *Field) {
	tCell := make([][]Cell, length*width)
	retField.Grid = &tCell
	retField.Length = length
	retField.Width = width
	return
}

// field.go
package main

import (
	"errors"
	"fmt"
)

type State int

const ( //Those are all the states of the cell in the battledfield
	empty State = iota
	full
	miss
	hit
)

var NoRep = errors.New("Can't shoot at things repeatedly")
var WrStatErr = errors.New("Wat? This can not be!")

func (st *State) change() error {
	switch *st { //My first switch, ever!
	case empty:
		*st = miss
	case full:
		*st = hit
	case miss, hit:
		return NoRep
	default:
		return WrStatErr
	}
	return nil
}

type Cell struct {
	state  State
	weight float64 //Probability of the ship inside, for AI to target. Will be tricky
}

type Field struct {
	Grid          *[][]Cell
	Length, Width int //Creating field based on rules
}

func newField(length int, width int) (retField *Field) {
	tCell := make([][]Cell, length*width)
	retField.Grid = &tCell
	retField.Length = length //just in case
	retField.Width = width
	return
}

func (Fi *Field) putShips(shs map[string]ship, act Actor) error {
	for name, params := range shs {
		fmt.Println(name)
		for i := 0; i < params.Numbers; i++ {
			if err := act.Putter(Fi, params.Fatness); err != nil {
				return err
			}
		}
	}

	return nil
}

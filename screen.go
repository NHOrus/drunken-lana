// screen
package main

import (
	//	"fmt"
	"errors"
	termbox "github.com/nsf/termbox-go"
)

func drawfield(lField <-chan Field, rField <-chan Field) error {
	var lFl, rFl Field

	//Straight up copy from docs about how initialize our little image
	if err := termbox.Init(); err != nil {
		return err
	}
	defer termbox.Close()

	//	termBuf := termbox.CellBuffer()
	termSizW, termSizH := termbox.Size()

	//We are getting field parameters. Possibly need to read them from rules, but...
	lFl = <-lField
	rFl = <-rField
	//Sanity checks
	if (lFl.Length != rFl.Length) && (lFl.Width != rFl.Width) {
		return errors.New("Can't do mismatched fields yet")
	}

	sFlen := lFl.Length + rFl.Length + 1
	sFwid := lFl.Width + rFl.Width + 3
	//one field in Heigth, two fields in Width, plus line for identifications of rows and columns

	//And more sanity checks
	if (termSizW < sFwid) && (termSizH < sFlen) {
		return errors.New("Terminal too small")
	}

	return nil //Yay!
}

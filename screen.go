// screen
package main

import (
	//	"fmt"
	"errors"
	termbox "github.com/nsf/termbox-go"
)

func drawfield(lField <-chan Field, rField <-chan Field) error {
	var lFl, rFl Field

	if err := termbox.Init(); err != nil {
		return err
	}
	defer termbox.Close()

	//	termBuf := termbox.CellBuffer()
	termSizW, termSizH := termbox.Size()

	lFl = <-lField
	rFl = <-rField

	if (lFl.Length != rFl.Length) && (lFl.Width != rFl.Width) {
		return errors.New("Can't do mismatched fields yet")
	}

	sFlen := lFl.Length + rFl.Length + 1
	sFwid := lFl.Width + rFl.Width + 3

	if (termSizW < sFwid) && (termSizH < sFlen) {
		return errors.New("Terminal too small")
	}

	return nil
}

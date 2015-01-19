// screen
package main

import (
	"errors"
	"fmt"
	termbox "github.com/nsf/termbox-go"
	"time"
)

func drawfield(lField <-chan Field, rField <-chan Field) error {
	var lFl, rFl Field
	fmt.Println("We are in")
	//Straight up copy from docs about how initialize our little image
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	//fmt.Println("Initialized")
	termSizW, termSizH := termbox.Size()

	//fmt.Println(termSizH, termSizW)
	//We are getting field parameters. Possibly need to read them from rules, but...
	lFl = <-lField
	rFl = <-rField

	//Sanity checks
	if (lFl.Fheight != rFl.Fheight) && (lFl.Fwidth != rFl.Fwidth) {
		return errors.New("Can't do mismatched fields yet")
	}

	sFhei := int(lFl.Fheight + 1)
	sFwid := int(lFl.Fwidth + rFl.Fwidth + 3)
	//one field in Heigth, two fields in Width, plus line for identifications of rows and columns

	//And more sanity checks
	if (termSizW < sFwid) && (termSizH < sFhei) {
		return errors.New("Terminal too small")
	}

	for varW := (termSizW - sFwid); varW < termSizW; varW++ {
		for varH := (termSizH - sFhei); varH < termSizH; varH++ {
			termbox.SetCell(varW, varH, 'x', termbox.ColorBlack, termbox.ColorWhite)

		}
	}

	termbox.Flush()
	time.Sleep(3 * time.Second)
	return nil //Yay!
}

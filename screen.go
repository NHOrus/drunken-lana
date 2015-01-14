// screen
package main

import (
	//	"fmt"
	termb "github.com/nsf/termbox-go"
)

func drawfield() error {
	if err := termb.Init(); err != nil {
		return err
	}
	return nil
}

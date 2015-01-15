// actors
package main

type Actor interface {
	Putter(field *Field, fat int) error //Deploy the ships!
	Hitter(field *Field) error          //Shoot! Shoot!
	Presenter()                         //Currently does nothing, ought to give enemy field as actor sees it.
}

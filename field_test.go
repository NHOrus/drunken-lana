// field_test
package main

import (
	"testing"
)

func TestChange(t *testing.T) {

	var WrState State = 10
	cases := []struct {
		in, want State
		correrr  error
	}{
		{empty, miss, nil}, //Normal change of state
		{full, hit, nil},
		{hit, hit, NoRep}, //Abmormal change of state - testing for errors
		{miss, miss, NoRep},
		{WrState, WrState, WrStatErr},
	}
	var mod State
	for _, c := range cases {
		mod = c.in
		err := mod.change()
		if (err != c.correrr) && (mod != c.want) {
			t.Errorf("Oops")
		}
	}
}

package cmd

import (
	"testing"
)

func TestChoose(T *testing.T) {
	// testing the arguments
	var args []string
	args = append(args, "d")
	got := chooseJokes(args)
	want := false
	if got != want {
		T.Errorf("TEST CASE FAILED. Got %v, Wanted %v", got, want)
	} else {
		T.Logf("PASSED.")
	}

	var args1 []string
	args1 = append(args1, "1", "2")
	got1 := chooseJokes(args1)
	want1 := false
	if got1 != want1 {
		T.Errorf("TEST CASE FAILED. Got %v, Wanted %v", got1, want1)
	} else {
		T.Logf("PASSED.Got %v, Wanted %v", got1, want1)
	}

	var args2 []string
	args2 = append(args2, "def", "50", "a123")
	got2 := chooseJokes(args2)
	want2 := false
	if got2 != want2 {
		T.Errorf("TEST CASE FAILED. Got %v, Wanted %v", got2, want2)
	} else {
		T.Logf("PASSED.Got %v, Wanted %v", got2, want2)
	}
}

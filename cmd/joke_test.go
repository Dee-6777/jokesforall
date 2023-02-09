package cmd

import (
	"testing"
)

func TestJoke(T *testing.T) {
	// testing if we're getting an empty string
	got := len(GetRandomJoke()) > 0
	want := true
	if got != want {
		T.Errorf("TEST CASE FAILED. Got %v, Wanted %v", got, want)
	} else {
		T.Logf("PASSED.")
	}

	// testing if we're getting an empty slice of bytes
	url := "https://icanhazdadjoke.com/"
	responseBytes := GetJokeData(url)
	var got_ bool
	if len(responseBytes) > 0 {
		got_ = true
	} else {
		got_ = false
	}
	want_ := true
	if got_ != want_ {
		T.Errorf("TEST CASE FAILED. Got %v, Wanted %v", got_, want_)
	} else {
		T.Logf("PASSED.")
	}
}

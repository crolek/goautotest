package main

//test

import (
	"testing"
)

func TestSample(t *testing.T) {
	// Do nothing
}

func TestFailure(t *testing.T) {
	//the line below will fail and the command line will beep at you :)
	t.Fail()
}

package main

//test

import (
	"log"
	"testing"
)

func TestLog(t *testing.T) {
	// Do nothing
	log.Println("First test log.Println statement")
}

func TestFailure(t *testing.T) {
	//the line below will fail and the command line will beep at you :)
	log.Println("A log.Println() statement")
	t.Log("Log before fail")
	t.Fail()
}

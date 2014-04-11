package main

import (
	"bytes"
	"fmt"
	"github.com/howeyc/fsnotify"
	"os"
	"os/exec"
	"strings"
	"time"
)

func startGoTest(doneChan chan bool) {
	var output bytes.Buffer
	fmt.Println("Running tests...")

	args := append([]string{"test"}, os.Args[1:]...)
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//snagging the cmd output.
	cmd.Stdout = &output
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	//checking to see if any unit tests failed, if so do the windows cmd beep
	if strings.Contains(output.String(), "--- FAIL") {
		fmt.Print("\x07") //the lovely console beep sound :D
	}

	//dislay out the unit test results
	fmt.Println(output.String())

	fmt.Println()
	fmt.Println("waiting...")
	doneChan <- true
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = watcher.Watch(wd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer watcher.Close()

	ignore := false
	doneChan := make(chan bool)
	readyChan := make(chan bool)

	//initial waiting message
	fmt.Println("waiting...")
	for {
		select {
		case ev := <-watcher.Event:
			if strings.HasSuffix(ev.Name, ".go") && !ignore {
				ignore = true
				go startGoTest(doneChan)
			}

		case err := <-watcher.Error:
			fmt.Println(err)

		case <-doneChan:
			time.AfterFunc(1500*time.Millisecond, func() {
				readyChan <- true
			})

		case <-readyChan:
			ignore = false
		}
	}

}

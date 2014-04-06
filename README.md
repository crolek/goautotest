##GoAutoTest
Simple tool to automatically run "go test" when a change is made to .go files. This is forked off of [https://github.com/ryanslade/goautotest/]().

This fork will throw a windows cmd `beep` at you when tests fail.

###INSTALL
```
go get github.com/crolek/goautotest
```

###TO RUN
Run "goautotest" in the directory of your project.
All the normal arguments that can be passed to "go test" can also be passed to "goautotest"

If you want to create an executable to move it to your project do:

`go build goautotest.go`

Then just copy/paste the `goautotest.exe` into your project's unit test directory.

###Warning
If the projct you are using this on has build errors it doesn't properly report them. Sorry.

###Todo
* Fix the build reporting bug
* Allow you to point this at a directory and let it run
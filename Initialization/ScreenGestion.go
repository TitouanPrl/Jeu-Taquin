package Initialization

import (
	"os"
	"os/exec"
	"runtime"
)

var clearTerminal map[string]func()

func init() {
	clearTerminal = make(map[string]func())

	/* Defining a clear method for multiple OS */
	clearTerminal["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
	clearTerminal["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}
}

func CallClearTerminal() {
	/* Checking the OS in use */
	value, ok := clearTerminal[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clearTerminal terminal screen :(")
	}
}

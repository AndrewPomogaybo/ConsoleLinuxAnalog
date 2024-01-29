package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help":
			fmt.Print(ClearHelp())
			return
		}
	}
	Default_clear()
}

func ClearHelp() string {
	return "clear: invalid option -- '-'\nUsage: clear [options]\n\nOptions:\n  -T TERM     use this instead of $TERM\n  -V          print curses-version\n  -x          do not try to clear scrollback\n"
}

func Default_clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

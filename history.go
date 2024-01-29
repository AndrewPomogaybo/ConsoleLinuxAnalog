package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Print(histopry_Help())
		return
	}
	fmt.Print(History("history"))
}

func History(command string) string {
	cmd := exec.Command(command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	return string(output)
}

func histopry_Help() string {
	return "history: history [-c] [-d смещение] [n] или history -anrw [файл] или history -ps аргумент [аргумент...]\n    Display or manipulate the history list.\n    \n    Display the history list with line numbers, prefixing each modified\n    entry with a `*'.  An argument of N lists only the last N entries.\n    \n    Options:\n      -c\tclear the history list by deleting all of the entries\n      -d offset\tdelete the history entry at position OFFSET. Negative\n    \t\toffsets count back from the end of the history list\n    \n      -a\tappend history lines from this session to the history file\n      -n\tread all history lines not already read from the history file\n    \t\tand append them to the history list\n      -r\tread the history file and append the contents to the history\n    \t\tlist\n      -w\twrite the current history to the history file\n    \n      -p\tperform history expansion on each ARG and display the result\n    \t\twithout storing it in the history list\n      -s\tappend the ARGs to the history list as a single entry\n    \n    If FILENAME is given, it is used as the history file.  Otherwise,\n    if HISTFILE has a value, that is used, else ~/.bash_history.\n    \n    If the HISTTIMEFORMAT variable is set and not null, its value is used\n    as a format string for strftime(3) to print the time stamp associated\n    with each displayed history entry.  No time stamps are printed otherwise.\n    \n    Exit Status:\n    Returns success unless an invalid option is given or an error occurs.\n"
}

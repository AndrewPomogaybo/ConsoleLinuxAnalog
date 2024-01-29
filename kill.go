package main

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <pid>\n", os.Args[0])
		return
	}

	if os.Args[1] == "--help" {
		fmt.Print(kill_Help())
		return
	}
	pid, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Invalid pid: %s\n", os.Args[1])
		return
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Printf("Failed to find process: %s\n", err)
		return
	}

	err = process.Signal(syscall.SIGKILL)
	if err != nil {
		fmt.Printf("Failed to kill process: %s\n", err)
		return
	}

	fmt.Printf("Process with PID %d killed\n", pid)

}

func kill_Help() string {
	return "kill: kill [-s назв_сигнала | -n номер_сигнала | -назв_сигнала] ид_процесса | назв_задания] ... или kill -l [назв_сигнала]\n    Send a signal to a job.\n    \n    Send the processes identified by PID or JOBSPEC the signal named by\n    SIGSPEC or SIGNUM.  If neither SIGSPEC nor SIGNUM is present, then\n    SIGTERM is assumed.\n    \n    Options:\n      -s sig\tSIG is a signal name\n      -n sig\tSIG is a signal number\n      -l\tlist the signal names; if arguments follow `-l' they are\n    \t\tassumed to be signal numbers for which names should be listed\n      -L\tsynonym for -l\n    \n    Kill is a shell builtin for two reasons: it allows job IDs to be used\n    instead of process IDs, and allows processes to be killed if the limit\n    on processes that you can create is reached.\n    \n    Exit Status:\n    Returns success unless an invalid option is given or an error occurs.\n"
}

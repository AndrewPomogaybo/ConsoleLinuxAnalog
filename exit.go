package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help":
			fmt.Print(ExitHelp())
			return
		}
	}
	ExitConsole()
}

func ExitConsole() {
	cmd := exec.Command("pidof", "mate-terminal")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	strData := string(output)
	str := strings.TrimSpace(strData)
	intValue, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	er := syscall.Kill(intValue, syscall.SIGKILL)
	if er != nil {
		log.Fatal(er)
	}
}

func ExitHelp() string {
	return "exit: exit [n]\n    Выход из командного процессора.\n    \n    Закрывает командный процессор с состоянием N. Если N не указан,\n    состоянием выхода будет состояние последней выполненной команды.\n"
}

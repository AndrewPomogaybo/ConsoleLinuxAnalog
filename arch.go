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
			fmt.Print(ArchHelp())
			break
		}
	} else {
		fmt.Print(GetArch())
	}

}

func ArchHelp() string {
	return "Использование: arch [ПАРАМЕТР]…\nПечатает машинную архитектуру.\n\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/arch>\nили доступная локально: info '(coreutils) arch invocation'\n"
}

func GetArch() string {
	cmd := exec.Command("arch")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(output)
}

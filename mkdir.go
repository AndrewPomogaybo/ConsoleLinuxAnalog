package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		if os.Args[1] == "-p" {
			if strings.Contains(os.Args[2], "/") {
				err := os.MkdirAll(os.Args[2], 0755)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				length := len(os.Args)
				for i := 2; i < length; i++ {
					drName1 := os.Args[i]
					os.Mkdir(drName1, 0755)
				}
			}
		} else if os.Args[1] == "--help" {
			mkdir_Help()
		} else {
			length := len(os.Args)
			for i := 1; i < length; i++ {
				drName1 := os.Args[i]
				os.Mkdir(drName1, 0755)
			}
		}
	} else {
		drName := os.Args[1]
		os.Mkdir(drName, 0755)
	}
}

func mkdir_Help() string {
	return "Использование: mkdir [ПАРАМЕТР]… КАТАЛОГ…\nСоздаёт КАТАЛОГ(и), если он ещё не существует.\n\nАргументы, обязательные для длинных параметров, обязательны и для коротких.\n  -m, --mode=РЕЖИМ   установить права доступа к файлу (как в chmod),\n                     а не a=rwx - umask\n  -p, --parents      не выдавать ошибку, если существует; создавать\n                     родительские каталоги при необходимости\n  -v, --verbose      печатать сообщение о каждом созданном каталоге\n  -Z                 установить контекст безопасности SELinux\n                     каждого создаваемого каталога равным типу по умолчанию\n      --context[=CTX]  подобно -Z, или если указан CTX, то установить\n                     контекст безопасности SELinux\n                     или SMACK равным CTX\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/mkdir>\nили доступная локально: info '(coreutils) mkdir invocation'"
}

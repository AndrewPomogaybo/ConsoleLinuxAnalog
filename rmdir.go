package main

import (
	"fmt"
	"os"
)

func main() {
	options := os.Args[1:]

	switch os.Args[1] {
	case "--help":
		fmt.Print(rmdir_Help())
		return
	}

	for i := 0; i < len(options); i++ {
		err := os.Remove(options[i])
		if err != nil {
			fmt.Println(err)
		}
	}
}

func rmdir_Help() string {
	return "Использование: rmdir [ПАРАМЕТР]… КАТАЛОГ…\nУдаляет КАТАЛОГ(и), если они пусты.\n\n      --ignore-fail-on-non-empty\n                  игнорировать все ошибки, которые возникают из-за того, что\n                  каталог не пуст\n  -p, --parents   удалить КАТАЛОГ и его потомков. Например, «rmdir -p a/b/c»\n                  делает то же, что и «rmdir a/b/c a/b a».\n  -v, --verbose   выводить сообщение для каждого обработанного каталога\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/rmdir>\nили доступная локально: info '(coreutils) rmdir invocation'\n"
}

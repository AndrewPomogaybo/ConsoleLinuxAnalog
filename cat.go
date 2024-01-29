package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	switch os.Args[1] {
	case "-n":
		cat_n()
		return
	case "--help":
		fmt.Print(CatHelp())
		return
	}
	Default_cat()
}

func CatHelp() string {
	return "Использование: cat [ПАРАМЕТР]… [ФАЙЛ]…\nПечатает слияние ФАЙЛ(ов) на стандартный вывод.\n\nЕсли ФАЙЛ не задан или задан как -, читает стандартный ввод.\n\n  -A, --show-all           синоним -vET\n  -b, --number-nonblank    нумеровать непустые строки при выводе\n  -e                       подавляет -n, синоним -vE\n  -E, --show-ends          показывать $ в конце каждой строки\n  -n, --number             нумеровать все строки при выводе\n  -s, --squeeze-blank      выводить не более одной пустой строки при повторе\n  -t                       синоним -vT\n  -T, --show-tabs          показывать символы табуляции как ^I\n  -u                       (игнорируется)\n  -v, --show-nonprinting   использовать запись с ^ и M-, за исключением\n                           символов перевода строки и табуляции\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nПримеры:\n  cat f - g  Выводит содержимое f, потом стандартный ввод, потом g.\n  cat        Копирует стандартный ввод на стандартный вывод.\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/cat>\nили доступная локально: info '(coreutils) cat invocation'\n"
}

func cat_n() {
	file, err := os.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	strfile := string(file)
	lines := strings.Split(strfile, "\n")
	for c, line := range lines {
		linenum := c + 1
		fmt.Printf("%d %s \n", linenum, line)
	}
}

func Default_cat() {
	for _, file := range os.Args[1:] {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))
	}
}

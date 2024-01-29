package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case "-n":
			Head_Option_MinusN()
			return
		case "-c":
			Head_Option_MinusC()
			return
		case "--help":
			fmt.Print(HeadHelp())
			return
		}

		for i := 1; i < len(os.Args); i++ {
			fmt.Println("==> " + os.Args[i] + " <==")
			file, err := os.ReadFile(os.Args[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			strfile := string(file)
			lines := strings.Split(strfile, "\n")
			for i := 0; i < 10; i++ {
				fmt.Println(lines[i])
			}
			fmt.Println("\n")
		}

	} else {
		Default_Head()
	}
}

func HeadHelp() string {
	return "спользование: head [ПАРАМЕТР]… [ФАЙЛ]…\nПечатает первые 10 строк каждого ФАЙЛА на стандартный вывод.\nЕсли задано несколько ФАЙЛОВ, сначала печатает заголовок с именем файла.\n\nЕсли ФАЙЛ не задан или задан как -, читает стандартный ввод.\n\nАргументы, обязательные для длинных параметров, обязательны и для коротких.\n  -c, --bytes=[-]K         напечатать первые K байт каждого файла;\n                           если перед K стоит «-», напечатать все, кроме K\n                           последних байт каждого файла\n  -n, --lines=[-]K         напечатать K строк каждого файла, а не первые 10;\n                           если перед K стоит «-», напечатать все, кроме K\n                           последних строк каждого файла\n  -q, --quiet, --silent    не печатать заголовки с именами файлов\n  -v, --verbose            всегда печатать заголовки с именами файлов\n  -z, --zero-terminated    разделитель строк NUL, а не символ\n                           новой строки\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nПосле K может стоять один из умножающий суффиксов:\nb 512, kB 1000, K 1024, MB 1000*1000, M 1024*1024,\nGB 1000*1000*1000, G 1024*1024*1024 и так далее для T, P, E, Z, Y.\nТакже можно использовать двоичные префиксы: KiB=K, MiB=M и так далее.\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/head>\nили доступная локально: info '(coreutils) head invocation'\n"
}

func Head_Option_MinusN() {
	file, err := os.ReadFile(os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
	strfile := string(file)
	lines := strings.Split(strfile, "\n")
	countline, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Вы не указали количество строк!")
		return
	}
	for i := 0; i < countline; i++ {
		fmt.Println(lines[i])
	}
}

func Head_Option_MinusC() {
	content, err := os.ReadFile(os.Args[3])
	if err != nil {
		fmt.Print("Ошибка при чтении файла", err)
	}
	bytesToPrint, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Print("Вы не указали количество байт")
	}
	if len(content) < bytesToPrint {
		bytesToPrint = len(content)
	}
	fmt.Println(string(content[:bytesToPrint]))
}

func Default_Head() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	strfile := string(file)
	lines := strings.Split(strfile, "\n")
	for i := 0; i < 10; i++ {
		fmt.Println(lines[i])
	}
}

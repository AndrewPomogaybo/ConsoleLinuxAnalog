package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if len(os.Args) > 2 {
		if os.Args[1] == "echo" {
			text := os.Args[2]
			fmt.Print(LineCounterInString(text), "   ", WordsCounterInString(text), "   ", GetStringSize(text), "\n")
			return
		} else {
			fmt.Print("invalid syntax")
			return
		}
	}

	file := os.Args[1]
	switch file {
	case "--help":
		fmt.Print(wc_Help())
		return
	}
	fmt.Println(WordsCounter(file), "  ", LinesCounter(file), "   ", getFileSize(file), "   ")
}

func WordsCounter(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	wordCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wordCount
}

func LinesCounter(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lineCount
}

func getFileSize(filePath string) int64 {
	file, err := os.Stat(filePath)
	if err != nil {
		return 0
	}
	size := file.Size()
	return size
}

func wc_Help() string {
	return "Использование: wc [ПАРАМЕТР]… [ФАЙЛ]…\n       или:    wc [ПАРАМЕТР]… --files0-from=Ф\nПечатает число символов новой строки, слов и байт для каждого ФАЙЛА и\nитоговую строку, если было задано несколько ФАЙЛОВ. Словом считается\nпоследовательность символов ненулевой длины, отделённая пробельным символом.\n\nЕсли ФАЙЛ не задан или задан как -, читает стандартный ввод.\n\nДля выбора выводимых счётчиков используются следующие параметры\n(счётчики всегда выводятся в таком порядке: символы новой строки,\nслова, символы, байты, максимальная длина строки):\n  -c, --bytes            напечатать количество байт\n  -m, --chars            напечатать количество символов\n  -l, --lines            напечатать количество новых строк\n      --files0-from=Ф    читать ввод из файлов, имена которых\n                         (завершённые нулем) перечислены в файле Ф;\n                         Если Ф равно -, то читать имена файлов из\n                         стандартного ввода\n  -L, --max-line-length  напечатать максимальную выводимую ширину\n  -w, --words            напечатать количество слов\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/wc>\nили доступная локально: info '(coreutils) wc invocation'\n"
}

func WordsCounterInString(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	wordCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += len(words)
	}

	return wordCount
}

func LineCounterInString(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}

func GetStringSize(str string) int {
	return utf8.RuneCountInString(str)
}

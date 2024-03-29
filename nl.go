package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if os.Args[1] == "--help" {
		fmt.Print(nl_Help())
		return
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	str := string(file)
	lines := strings.Split(str, "\n")
	for c, line := range lines {
		numbers := c + 1
		fmt.Printf("%d %s \n", numbers, line)
	}
}

func nl_Help() string {
	return "Использование: nl [ПАРАМЕТР]… [ФАЙЛ]…\nПечатает каждый ФАЙЛ на стандартный вывод, добавляя номера строк.\n\nЕсли ФАЙЛ не задан или задан как -, читает стандартный ввод.\n\nАргументы, обязательные для длинных параметров, обязательны и для коротких.\n  -b, --body-numbering=СТИЛЬ      использовать СТИЛЬ нумерования строк тела\n  -d, --section-delimiter=СС      использовать СС как логический разделитель\n                                  страниц\n  -f, --footer-numbering=СТИЛЬ    использовать СТИЛЬ нумерования строк нижнего\n                                  колонтитула\n  -h, --header-numbering=СТИЛЬ    использовать СТИЛЬ нумерования строк верхнего\n                                  колонтитула\n  -i, --line-increment=ЧИСЛО      шаг увеличения номеров строк\n  -l, --join-blank-lines=ЧИСЛО    заданное ЧИСЛО пустых строк считать одной\n  -n, --number-format=ФОРМАТ      использовать ФОРМАТ для номеров строк\n  -p, --no-renumber               не начинать нумерацию заново для каждого\n                                  раздела\n  -s, --number-separator=СТРОКА   добавлять СТРОКУ после номера\n  -v, --starting-line-number=ЧИСЛО первый номер строки для каждого раздела\n  -w, --number-width=ЧИСЛО        использовать заданное ЧИСЛО столбцов для\n                                  номеров строк\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nПараметры по умолчанию: -bt -d'\\:' -fn -hn -i1 -l1 -n'rn' -s<TAB> -v1 -w6\n\nСС — это два символа, используемые для создания логического разделителя\nстраниц; при отсутствии второго используется «:».\n\nСТИЛЬ задается как один из следующих:\n\n  a         нумеровать все строки\n  t         нумеровать только непустые строки\n  n         не нумеровать строки\n  pБРВ      нумеровать только строки, часть которых совпадает с базовым\n              регулярным выражением БРВ\n\nФОРМАТ задается как один из следующих:\n\n  ln   выравнивать по левому краю, не выводить начальные нули\n  rn   выравнивать по правому краю, не выводить начальные нули\n  rz   выравнивать по правому краю, выводить начальные нули\n\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/nl>\nили доступная локально: info '(coreutils) nl invocation'\n"
}

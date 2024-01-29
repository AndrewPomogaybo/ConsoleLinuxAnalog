package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	switch os.Args[1] {
	case "--help":
		fmt.Print(touch_Help())
		return
	}

	options := os.Args[1:]
	for i := 0; i < len(options); i++ {
		if _, err := os.Stat(options[i]); os.IsNotExist(err) {
			file, err := os.Create(options[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			file.Close()
		} else {
			// Изменяем время модификации файла
			currentTime := time.Now()
			err := os.Chtimes(options[i], currentTime, currentTime)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func touch_Help() string {
	return "Использование: touch [ПАРАМЕТР]… ФАЙЛ…\nУстанавливает временные метки доступа и изменения каждого ФАЙЛА в текущее\nвремя.\n\nЕсли ФАЙЛ не существует, то создаётся пустой, если не указаны\nпараметры -c или -h.\n\nЕсли строка аргумента ФАЙЛ равна -, то это обрабатывается специальным образом\nи вызывает изменение временных меток файла, связанного со стандартным выводом.\n\nАргументы, обязательные для длинных параметров, обязательны и для коротких.\n  -a                     изменить только время доступа\n  -c, --no-create        не создавать файлы\n  -d, --date=СТРОКА      проанализировать СТРОКУ и использовать её вместо\n                         текущего времени\n  -f                     (игнорируется)\n  -h, --no-dereference   изменяет символьные ссылки, а не файлы, на которые\n                         они ссылаются (имеет смысл только на системах, где\n                         можно изменять владельца символьной ссылки)\n  -m                     изменять только время изменения\n  -r, --reference=ФАЙЛ   использовать время ФАЙЛА вместо текущего времени\n  -t ВРЕМЯ               использовать [[ВВ]ГГ]ММДДччмм[.сс] вместо\n                         текущего времени\n  --time=СЛОВО           изменить время, указанное СЛОВОм:\n                         значения access, atime, use эквивалентны -a\n                         значения modify, mtime эквивалентны -m\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nЗаметьте, что параметры -d и -t используют разные форматы даты и времени.\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/touch>\nили доступная локально: info '(coreutils) touch invocation'\n"
}

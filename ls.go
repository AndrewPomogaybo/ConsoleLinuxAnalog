package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"
)

func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help":
			fmt.Print(LsHelp())
			return
		case "-R":
			ls_R()
			return
		case "-a":
			ls_a()
			return
		case "-l":
			ls_l()
			return
		case "-r":
			ls_r()
			return
		case "-h":
			ls_h()
			return
		}
	}
	Defualt_ls()
}

func LsHelp() string {
	return "Использование: ls [ПАРАМЕТР]… [ФАЙЛ]…\nВыдаёт информацию о ФАЙЛАХ (по умолчанию о текущем каталоге).\nСортирует в алфавитном порядке, если не задан ни --sort, ни один из\nпараметров -cftuvSUX.\n\nАргументы, обязательные для длинных параметров, обязательны и для коротких.\n  -a, --all                  не скрывать файлы начинающиеся с .\n  -A, --almost-all           не выдавать подразумеваемые . и ..\n      --author               вместе с -l, печатать автора каждого файла\n  -b, --escape               печатать экранирующие последовательности\n                             в стиле С для не графических символов\n      --block-size=РАЗМЕР    использовать блоки указанного РАЗМЕРА; например,\n                             «---block-size=M»; см. формат РАЗМЕРА далее.\n  -B, --ignore-backups       не выдавать файлы, оканчивающиеся на ~\n  -c                         с -lt: сортировать и показывать по ctime (времени\n                             последнего изменения файла);\n                             с -l: показывать ctime и сортировать по имени;\n                             иначе: сортировать по ctime, сначала самые новые\n  -C                         выдавать список в несколько колонок\n      --color[=КОГДА]        расцвечивать вывод;\n                             КОГДА может быть «always» («всегда»,\n                             по умолчанию, если не указано), «auto» или\n                             «never». Подробней см. ниже\n  -d, --directory            выдавать имена каталогов, а не их содержимое\n  -D, --dired                генерировать вывод для режима Emacs dired\n  -f                         не сортировать, включает -aU,\n                             выключает -ls --color\n  -F, --classify             добавлять к элементам индикатор (один из */=>@|)\n      --file-type            аналогично, но не добавлять «*»\n      --format=СЛОВО         across -x, commas -m, horizontal -x, long -l,\n                             single-column -1, verbose -l, vertical -C\n      --full-time            синоним -l --time-style=full-iso\n  -g                         как -l, но не выводить владельца\n      --group-directories-first\n                             группировать каталоги раньше файлов;\n                             может дополняться параметром --sort,\n                             но любое использование\n                             --sort=none (-U) выключает группировку\n  -G, --no-group             в длинном формате не печатать имена групп\n  -h, --human-readable       c -l и/или -s печатать размеры в удобном для\n                             человека виде (например, 1K 234M 2G)\n      --si                   то же, но использовать степень 1000, а не 1024\n  -H, --dereference-command-line\n                             следовать по символьным ссылкам в командной\n                             строке\n      --dereference-command-line-symlink-to-dir\n                             следовать по всем символьным ссылкам в командной\n                             строке, которые указывают на каталог\n      --hide=ШАБЛОН          не показывать записи, соответствующие ШАБЛОНУ\n                             оболочки (отменяется параметрами -a или -A)\n      --hyperlink[=КОГДА]    имена файлов гиперссылок; КОГДА может быть\n                             «always» (по умолчанию, если не задано),\n                             «auto» или «never»\n      --indicator-style=СТИЛЬ добавлять к именам элементов индикатор с\n                             указанным СТИЛЕМ: none (по умолчанию),\n                             slash (-p), file-type (--file-type), classify (-F)\n  -i, --inode                показывать иноду каждого файла\n  -I, --ignore=ШАБЛОН        не показывать записи, соответствующие ШАБЛОНУ\n                             оболочки\n  -k, --kibibytes            по умолчанию использовать блоки по 1024 байта\n                             при показе использования диска;\n                             использовать только с -s и для итогового\n                             значения каталога\n  -l                         использовать широкий формат\n  -L, --dereference          показывая информацию для символьной ссылки,\n                             показывать информацию о файле, на который ссылка\n                             ссылается\n  -m                         выдавать список на всю ширину через запятую\n  -n, --numeric-uid-gid      как -l, но выдавать числовые UID и GID\n  -N, --literal              печатать имена как есть, без экранирования\n  -o                         как -l, но не выводить информацию о группе\n  -p,  --indicator-style=slash  добавлять индикатор / к каталогам\n  -q, --hide-control-chars   выводить ? вместо непечатных символов\n      --show-control-chars   показывать непечатные символы (по умолчанию,\n                             если программа не «ls» и вывод не на терминал).\n  -Q, --quote-name           заключать имя записей в кавычки\n      --quoting-style=ТИП    использовать ТИП заключения в кавычки:\n                             literal, locale, shell, shell-always,\n                             shell-escape, shell-escape-always, c, escape\n                             (заменяет переменную окружения QUOTING_STYLE)\n  -r, --reverse              обратный порядок сортировки\n  -R, --recursive            рекурсивно показывать каталоги\n  -s, --size                 печатать размер каждого файла в блоках\n  -S                         сортировать по размеру файлов, большие сначала\n      --sort=СЛОВО           сортировать по СЛОВУ, а не по имени: \n                             none (без сортировки) -U, size (размер) -S,\n                             time (время) -t, version (версия) -v,\n                             extension (расширение) -X\n                             \n      --time=СЛОВО           изменить значения по умолчанию, использующие\n                             время изменения;\n                             время доступа (-u): atime, access, use;\n                             время изменения (-c): ctime, status;\n                             время создания: birth, creation;\n                             с параметром -l показываемое время задаётся\n                             СЛОВОМ; если --sort=time, сортировать по СЛОВУ\n                             (сначала новые)\n      --time-style=TIME_STYLE  формат даты/времени с -l;\n                               смотрите описание TIME_STYLE ниже\n  -t                         сортировать по времени; смотрите --time\n  -T, --tabsize=РАЗМЕР       использовать табуляцию указанного РАЗМЕРА, а не 8\n  -u                         с -lt: сортировать и показывать время доступа;\n                             c -l: показывать время доступа и сортировать\n                             по имени; иначе сортировать по времени доступа,\n                             сначала самые новые\n  -U                         не сортировать; выводить в соответствии с\n                             физическим расположением в каталоге\n  -v                         сортировать по номерам (версии) в текстовом\n                             представлении\n  -w, --width=ЧИСЛО          устанавливает ширину вывода равной ЧИСЛО.\n                             0 означает отсутствие ограничения.\n  -x                         перечислять по строкам, а не по столбцам\n  -X                         сортировать по расширению в алфавитном порядке\n  -Z, --context              печатать любой контекст безопасности каждого\n                             файла\n  -1                         перечислять по одному файлу на строке.\n                             Символ «\\n» не учитывается, если указан -q или -b\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nРАЗМЕР задаётся целым числом и необязательной единицей измерения\n(пример: 10K это 10*1024). Единицы измерения:\nK, M, G, T, P, E, Z, Y (степень 1024) или KB, MB, … (степень 1000).\nТакже можно использовать двоичные префиксы: KiB=K, MiB=M и так далее.\n\nЗначением TIME_STYLE могут быть: full-iso, long-iso, iso, locale или +ФОРМАТ.\nЗначение ФОРМАТа как в date(1). Если значение ФОРМАТа равно\nФОРМАТ1<новая строка>ФОРМАТ2, то ФОРМАТ1 применяется не к последним файлам,\nа ФОРМАТ2 к новым. Если TIME_STYLE начинается с «posix-», то он применяется\nтолько для локалей отличных от POSIX. Также, используемый стиль по умолчанию\nзадаёт переменная окружения TIME_STYLE.\n\nИспользование цветов для различения типов файла по умолчанию выключено,\nэто же можно сделать и с помощью --color=never. С параметром --color=auto,\nls выдаёт цветовые коды только когда стандартный вывод подключён к терминалу.\nПеременная окружения LS_COLORS служит для изменения значений.\nДля её установки используйте команду dircolors.\n\nКоды выхода:\n 0  всё отлично,\n 1  небольшие проблемы (например, недоступен подкаталог),\n 2  серьёзная проблема (например, недоступен аргумент командной строки).\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/ls>\nили доступная локально: info '(coreutils) ls invocation'\n"
}

func Defualt_ls() {
	fileInfo, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range fileInfo {
		name := file.Name()
		if len(name) > 0 && name[0:1] != "." {
			fmt.Println(file.Name())
		}
	}
}

func ls_R() {
	dir := "."
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Ошибка при доступе к пути %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil // пропустить каталоги
		}
		fmt.Println(path)
		return nil
	})
	if err != nil {
		log.Printf("Ошибка при перечислении файлов: %v\n", err)
	}
}

func ls_a() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func ls_l() {
	fileInfos, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, fi := range fileInfos {
		info, err := fi.Info()
		if err != nil {
			log.Fatal(err)
		}

		fileName := info.Name()
		fileSize := info.Size()
		fileMode := info.Mode().String()
		linkCount := info.Sys().(*syscall.Stat_t).Nlink

		uid := info.Sys().(*syscall.Stat_t).Uid
		owner, err := user.LookupId(strconv.Itoa(int(uid)))
		if err != nil {
			log.Fatal(err)
		}

		ownerName := owner.Username
		gid := info.Sys().(*syscall.Stat_t).Gid
		group, err := user.LookupGroupId(strconv.Itoa(int(gid)))

		if err != nil {
			log.Fatal(err)
		}

		groupName := group.Name
		modTime := info.ModTime().Format("Jan _2 15:04")
		output := fmt.Sprintf("%s %3d %s %s %8d %s %s", fileMode, linkCount, ownerName, groupName, fileSize, modTime, fileName)
		fmt.Println(output)
	}
}

func ls_r() {
	dir := "."

	fileList := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Ошибка при обработке пути %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		log.Printf("Ошибка при выполнении команды ls -r: %v\n", err)
		return
	}

	for i := len(fileList) - 1; i >= 0; i-- {
		fmt.Println(fileList[i])
	}
}

func ls_h() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		info, err := os.Stat(file.Name())
		if err != nil {
			log.Fatal(err)
		}

		size := info.Size()
		var sizeStr string
		switch {
		case size > 1024*1024*1024:
			sizeStr = fmt.Sprintf("%.2fG", float64(size)/(1024*1024*1024))
		case size > 1024*1024:
			sizeStr = fmt.Sprintf("%.2fM", float64(size)/(1024*1024))
		case size > 1024:
			sizeStr = fmt.Sprintf("%.2fK", float64(size)/1024)
		default:
			sizeStr = strconv.Itoa(int(size)) + "B"
		}

		fmt.Printf("%s\t%s\n", sizeStr, file.Name())
	}
}

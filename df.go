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
			fmt.Print(df_Help())
			return
		case "-a":
			fmt.Print(df_a("df", "-a"))
			return
		}
	}
	fmt.Println(Default_df("df"))
}

func Default_df(command string) string {
	cmd := exec.Command(command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	return string(output)
}

func df_a(command, arg string) string {
	cmd := exec.Command(command, arg)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	return string(output)
}

func df_Help() string {
	return "Использование: df [ПАРАМЕТР]… [ФАЙЛ]…\nПоказывает сведения о файловой системе, на которой расположен каждый\nиз ФАЙЛОВ, или, по умолчанию, обо всех файловых системах.\n\nАргументы, обязательные для длинных параметров, обязательны и для коротких.\n  -a, --all             include pseudo, duplicate, inaccessible file systems\n  -B, --block-size=SIZE  scale sizes by SIZE before printing them; e.g.,\n                           '-BM' prints sizes in units of 1,048,576 bytes;\n                           see SIZE format below\n      --direct          show statistics for a file instead of mount point\n  -h, --human-readable  print sizes in powers of 1024 (e.g., 1023M)\n  -H, --si              print sizes in powers of 1000 (e.g., 1.1G)\n  -i, --inodes            вывести информацию об индексных дескрипторах, а не\n                          об использовании блоков\n  -k                      аналог --block-size=1K\n  -l, --local             перечислить только локальные файловые системы\n      --no-sync           не вызывать sync перед получением информации об\n                          использовании блоков (по умолчанию)\n      --output[=FIELD_LIST]\n                          при выводе использовать формат FIELD_LIST, или\n                          показать все поля, если значение FIELD_LIST\n                          не указано.\n  -P, --portability       выводить в формате POSIX\n      --sync              вызывать sync перед получением информации об\n                          использовании блоков\n      --total             не учитывать все маловажные для доступного\n                          пространства элементы и показать общий итог\n  -t, --type=ТИП          перечислить только файловые системы указанного ТИПА\n  -T, --print-type        выводить тип файловой системы\n  -x, --exclude-type=ТИП  исключить файловые системы указанного ТИПА\n  -v                      (игнорируется)\n      --help     показать эту справку и выйти\n      --version  показать информацию о версии и выйти\n\n\nПоказывает значения в единицах первого доступного РАЗМЕРА из --block-size,\nи переменных окружения DF_BLOCK_SIZE, BLOCK_SIZE и BLOCKSIZE.\nИначе, единицей по умолчанию является 1024 байта (или 512, если\nустановлена POSIXLY_CORRECT).\n\nРАЗМЕР задаётся целым числом и необязательной единицей измерения\n(пример: 10K это 10*1024). Единицы измерения:\nK, M, G, T, P, E, Z, Y (степень 1024) или KB, MB, … (степень 1000).\nТакже можно использовать двоичные префиксы: KiB=K, MiB=M и так далее.\n\nЗначение FIELD_LIST представляет собой список включаемых колонок,\nперечисляемых через запятую. Возможные имена полей:\n«source», «fstype», «itotal», «iused», «iavail», «ipcent»,\n«size», «used», «avail», «pcent», «file» и «target» (см. также страницу info).\n\nСтраница справки по GNU coreutils: <https://www.gnu.org/software/coreutils/>\nОб ошибках в переводе сообщений сообщайте по адресу <https://translationproject.org/team/ru.html>\nПолная документация: <https://www.gnu.org/software/coreutils/df>\nили доступная локально: info '(coreutils) df invocation'\n"
}

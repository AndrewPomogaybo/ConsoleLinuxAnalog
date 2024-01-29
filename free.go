package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "--help":
			fmt.Print(FreeHelp())
		}
	}
	fmt.Print(GetFree())
}

func FreeHelp() string {
	return "exit: exit [n]\n    Выход из командного процессора.\n    \n    Закрывает командный процессор с состоянием N. Если N не указан,\n    состоянием выхода будет состояние последней выполненной команды.\n[user@localhost ConsoleLinux]$ go build free.go\n[user@localhost ConsoleLinux]$ ./free\nTotal memory: 62424 bytes\n[user@localhost ConsoleLinux]$ free --help\n\nUsage:\n free [options]\n\nOptions:\n -b, --bytes         show output in bytes\n     --kilo          show output in kilobytes\n     --mega          show output in megabytes\n     --giga          show output in gigabytes\n     --tera          show output in terabytes\n     --peta          show output in petabytes\n -k, --kibi          show output in kibibytes\n -m, --mebi          show output in mebibytes\n -g, --gibi          show output in gibibytes\n     --tebi          show output in tebibytes\n     --pebi          show output in pebibytes\n -h, --human         show human-readable output\n     --si            use powers of 1000 not 1024\n -l, --lohi          show detailed low and high memory statistics\n -t, --total         show total for RAM + swap\n -s N, --seconds N   repeat printing every N seconds\n -c N, --count N     repeat printing N times, then exit\n -w, --wide          wide output\n\n     --help     display this help and exit\n -V, --version  output version information and exit\n\nFor more details see free(1).\n"
}

func GetFree() string {
	cmd := exec.Command("free")
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatal(err)

	}
	return string(stdout)
}

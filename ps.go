package main

import (
	"fmt"
	"os"
)

func main() {
	processes, err := os.ReadDir("/proc")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, process := range processes {
		if process.IsDir() {
			pid := process.Name()
			fmt.Printf("Имя процесса: %s, PID: %s\n", getProcessName(pid), pid)
		}
	}
}

func getProcessName(pid string) string {
	filePath := fmt.Sprintf("/proc/%s/comm", pid)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "Не удалось получить имя"
	}
	return string(data)
}

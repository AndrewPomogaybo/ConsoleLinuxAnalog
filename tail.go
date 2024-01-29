package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		numlines, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		file, err := os.ReadFile(os.Args[3])
		if err != nil {
			fmt.Println(err)
			return
		}
		str := string(file)
		lines := strings.Split(str, "\n")

		for _, line := range lines[len(lines)-numlines:] {
			fmt.Println(line)
		}
	} else {
		file, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		str := string(file)
		lines := strings.Split(str, "\n")
		for _, line := range lines[len(lines)-10:] {
			fmt.Println(line)
		}
	}
}

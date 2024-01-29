package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	com := exec.Command("file", os.Args[1])
	out, err := com.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}

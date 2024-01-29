package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: hexdump <filename>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	err = hexdump(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

func hexdump(file *os.File) error {
	buffer := make([]byte, 16)
	offset := 0

	for {
		numRead, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		fmt.Printf("%08x  ", offset)

		ascii := ""
		for i := 0; i < 16; i++ {
			if i < numRead {
				fmt.Printf("%02x ", buffer[i])
				if buffer[i] >= 32 && buffer[i] <= 126 {
					ascii += string(buffer[i])
				} else {
					ascii += "."
				}
			} else {
				fmt.Print("   ")
				ascii += " "
			}
			if i == 7 {
				fmt.Print(" ")
				ascii += " "
			}
		}

		fmt.Println(" |" + ascii + "|")
		offset += 16
	}

	return nil
}

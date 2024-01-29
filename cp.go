package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	sourceFilePath := os.Args[1]
	destinationFolder := os.Args[2]

	if isDirectory(destinationFolder) {
		err := copyFileIntoDir(sourceFilePath, destinationFolder)
		if err != nil {
			fmt.Println("Error copying file:", err)
			return
		}
		err = os.MkdirAll(destinationFolder, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating destination folder:", err)
			return
		}

		fmt.Printf("Файл '%s' успешно скопирован в папку '%s'\n", sourceFilePath, destinationFolder)
	} else {
		err := copyFileIntoFile(sourceFilePath, destinationFolder)
		if err != nil {
			fmt.Println("Error copying file:", err)
			return
		}

		fmt.Printf("Файл '%s' успешно скопирован в файл '%s'\n", sourceFilePath, destinationFolder)
	}

}

func copyFileIntoDir(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(filepath.Join(dest, filepath.Base(src)))
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func copyFileIntoFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	if !sourceFileInfo.Mode().IsRegular() {
		return fmt.Errorf("source file is not a regular file")
	}

	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func isDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}

package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	options := os.Args[1]

	switch options {
	case "-c":
		tar_c(os.Args[2], os.Args[3])
		return
	}

	Default_tar()
}

func tar_c(archiveName string, sourceDir string) {
	archiveFile, err := os.Create(archiveName)
	if err != nil {
		log.Fatal(err)
	}
	defer archiveFile.Close()

	tarWriter := tar.NewWriter(archiveFile)
	defer tarWriter.Close()

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			log.Fatal(err)
		}

		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			log.Fatal(err)
		}

		header.Name = relPath

		err = tarWriter.WriteHeader(header)
		if err != nil {
			log.Fatal(err)
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			_, err = io.Copy(tarWriter, file)
			if err != nil {
				log.Fatal(err)
			}
		}

		return nil
	}

	err = filepath.Walk(sourceDir, walkFunc)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Архивация завершена.")
}

func Default_tar() {
	files := os.Args[2:]

	err := createTarFile(os.Args[1], files)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Archive created successfully")
}

func addFileToTar(writer *tar.Writer, fileInfo os.FileInfo, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	header := &tar.Header{
		Name:    filePath,
		Size:    fileInfo.Size(),
		Mode:    int64(fileInfo.Mode()),
		ModTime: fileInfo.ModTime(),
	}

	if err := writer.WriteHeader(header); err != nil {
		return err
	}

	if _, err := io.Copy(writer, file); err != nil {
		return err
	}

	return nil
}

func createTarFile(tarFilePath string, files []string) error {
	tarFile, err := os.Create(tarFilePath)
	if err != nil {
		return err
	}
	defer tarFile.Close()

	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()

	for _, filePath := range files {
		err = filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			return addFileToTar(tarWriter, info, path)
		})

		if err != nil {
			return err
		}
	}

	return nil
}

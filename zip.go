package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	switch os.Args[1] {
	case "--help":
		fmt.Print(ZipHelp())
		return
	}

	files := os.Args[2:]
	archiveName := os.Args[1]

	err := zipFiles(archiveName, files)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Archive created:", archiveName)
}

func zipFiles(filename string, files []string) error {
	zipfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()

	for _, file := range files {
		if err := addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	wr, err := zipWriter.Create(filename)
	if err != nil {
		return err
	}

	if _, err = io.Copy(wr, file); err != nil {
		return err
	}

	return nil
}

func ZipHelp() string {
	return "Copyright (c) 1990-2008 Info-ZIP - Type 'zip \"-L\"' for software license.\nZip 3.0 (July 5th 2008). Usage:\nzip [-options] [-b path] [-t mmddyyyy] [-n suffixes] [zipfile list] [-xi list]\n  The default action is to add or replace zipfile entries from list, which\n  can include the special name - to compress standard input.\n  If zipfile and list are omitted, zip compresses stdin to stdout.\n  -f   freshen: only changed files  -u   update: only changed or new files\n  -d   delete entries in zipfile    -m   move into zipfile (delete OS files)\n  -r   recurse into directories     -j   junk (don't record) directory names\n  -0   store only                   -l   convert LF to CR LF (-ll CR LF to LF)\n  -1   compress faster              -9   compress better\n  -q   quiet operation              -v   verbose operation/print version info\n  -c   add one-line comments        -z   add zipfile comment\n  -@   read names from stdin        -o   make zipfile as old as latest entry\n  -x   exclude the following names  -i   include only the following names\n  -F   fix zipfile (-FF try harder) -D   do not add directory entries\n  -A   adjust self-extracting exe   -J   junk zipfile prefix (unzipsfx)\n  -T   test zipfile integrity       -X   eXclude eXtra file attributes\n  -y   store symbolic links as the link instead of the referenced file\n  -e   encrypt                      -n   don't compress these suffixes\n  -h2  show more help\n"
}

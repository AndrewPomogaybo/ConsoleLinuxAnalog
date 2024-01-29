package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	option := os.Args[1]

	switch option {
	case "--help":
		fmt.Print(UnzipHelp())
		return
	}

	err := unzipSource(os.Args[1], "")
	if err != nil {
		log.Fatal(err)
	}
}

func unzipSource(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	destination, err = filepath.Abs(destination)
	if err != nil {
		return err
	}

	for _, f := range reader.File {
		err := unzipFile(f, destination)
		if err != nil {
			return err
		}
	}

	return nil
}

func unzipFile(f *zip.File, destination string) error {

	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}

	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	zippedFile, err := f.Open()
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		return err
	}
	return nil
}

func UnzipHelp() string {
	return "UnZip 6.00 of 20 April 2009, by Info-ZIP.  Maintained by C. Spieler.  Send\nbug reports using http://www.info-zip.org/zip-bug.html; see README for details.\n\nUsage: unzip [-Z] [-opts[modifiers]] file[.zip] [list] [-x xlist] [-d exdir]\n  Default action is to extract files in list, except those in xlist, to exdir;\n  file[.zip] may be a wildcard.  -Z => ZipInfo mode (\"unzip -Z\" for usage).\n\n  -p  extract files to pipe, no messages     -l  list files (short format)\n  -f  freshen existing files, create none    -t  test compressed archive data\n  -u  update files, create if necessary      -z  display archive comment only\n  -v  list verbosely/show version info       -T  timestamp archive to latest\n  -x  exclude files that follow (in xlist)   -d  extract files into exdir\nmodifiers:\n  -n  never overwrite existing files         -q  quiet mode (-qq => quieter)\n  -o  overwrite files WITHOUT prompting      -a  auto-convert any text files\n  -j  junk paths (do not make directories)   -aa treat ALL files as text\n  -U  use escapes for all non-ASCII Unicode  -UU ignore any Unicode fields\n  -C  match filenames case-insensitively     -L  make (some) names lowercase\n  -X  restore UID/GID info                   -V  retain VMS version numbers\n  -K  keep setuid/setgid/tacky permissions   -M  pipe through \"more\" pager\n  -O CHARSET  specify a character encoding for DOS, Windows and OS/2 archives\n  -I CHARSET  specify a character encoding for UNIX and other archives\n\nSee \"unzip -hh\" or unzip.txt for more help.  Examples:\n  unzip data1 -x joe   => extract all files except joe from zipfile data1.zip\n  unzip -p foo | more  => send contents of foo.zip via pipe into program more\n  unzip -fo foo ReadMe => quietly replace existing ReadMe if archive file newer"
}

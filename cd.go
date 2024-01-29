package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	switch os.Args[1] {
	case "--help":
		fmt.Print(CdHelp())
		return
	}

	dir := os.Args[1]
	err := os.Chdir(dir)
	if err != nil {
		fmt.Printf("Failed to change directory: %s\n", err.Error())
		return
	}

	pwd, err := filepath.Abs(dir)
	if err != nil {
		fmt.Printf("Failed to get absolute path: %s\n", err.Error())
		return
	}

	fmt.Printf("Changed directory to: %s\n", pwd)
}

func CdHelp() string {
	return "cd: cd [-L|[-P [-e]] [-@]] [каталог]\n    Change the shell working directory.\n    \n    Change the current directory to DIR.  The default DIR is the value of the\n    HOME shell variable. If DIR is \"-\", it is converted to $OLDPWD.\n    \n    The variable CDPATH defines the search path for the directory containing\n    DIR.  Alternative directory names in CDPATH are separated by a colon (:).\n    A null directory name is the same as the current directory.  If DIR begins\n    with a slash (/), then CDPATH is not used.\n    \n    If the directory is not found, and the shell option `cdable_vars' is set,\n    the word is assumed to be  a variable name.  If that variable has a value,\n    its value is used for DIR.\n    \n    Options:\n      -L\tforce symbolic links to be followed: resolve symbolic\n    \t\tlinks in DIR after processing instances of `..'\n      -P\tuse the physical directory structure without following\n    \t\tsymbolic links: resolve symbolic links in DIR before\n    \t\tprocessing instances of `..'\n      -e\tif the -P option is supplied, and the current working\n    \t\tdirectory cannot be determined successfully, exit with\n    \t\ta non-zero status\n      -@\ton systems that support it, present a file with extended\n    \t\tattributes as a directory containing the file attributes\n    \n    The default is to follow symbolic links, as if `-L' were specified.\n    `..' is processed by removing the immediately previous pathname component\n    back to a slash or the beginning of DIR.\n    \n    Exit Status:\n    Returns 0 if the directory is changed, and if $PWD is set successfully when\n    -P is used; non-zero otherwise.\n"
}

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	if len(os.Args) == 1 {
		for i := 0; i < 20; i++ {
			fmt.Print(GenerateDefaultPwd("8"), " ", GenerateDefaultPwd("8"), " ", GenerateDefaultPwd("8"), " ", GenerateDefaultPwd("8"), " ", GenerateDefaultPwd("8"), " ", GenerateDefaultPwd("8"), " ", GenerateDefaultPwd("8"), " ", GenerateDefaultPwd("8"), "\n")
		}
		return
	}

	if len(os.Args) > 2 && os.Args[1] == "-1" {
		switch os.Args[2] {
		case "-s":
			fmt.Print(GeneratePwd("8"), "\n")
			return
		}
	}

	switch os.Args[1] {
	case "-y":
		for i := 0; i < 20; i++ {
			fmt.Print(GeneratePwd("8"), " ", GeneratePwd("8"), " ", GeneratePwd("8"), " ", GenerateDefaultPwd("8"), " ", GeneratePwd("8"), " ", GeneratePwd("8"), " ", GeneratePwd("8"), " ", GenerateDefaultPwd("8"), "\n")
		}
		return
	case "-1":
		fmt.Print(GenerateDefaultPwd("8"), "\n")
		return
	case "--help":
		fmt.Print(pwgen_Help())
		return
	case "-n":
		if len(os.Args) > 2 {
			length := os.Args[2]
			for i := 0; i < 20; i++ {
				fmt.Print(GeneratePwd(length), " ", GeneratePwd(length), " ", GeneratePwd(length), " ", GeneratePwd(length), " ", GeneratePwd(length), " ", GeneratePwd(length), " ", GeneratePwd(length), " ", GeneratePwd(length), "\n")
			}
			return
		}
	}
}

func GeneratePwd(length string) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	specials := "~=+%^*/()[]{}/!@#$?|"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits + specials
	intLen, err := strconv.Atoi(length)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, intLen)
	buf[0] = digits[rand.Intn(len(digits))]
	buf[1] = specials[rand.Intn(len(specials))]
	for i := 2; i < intLen; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)
	return str
}

func pwgen_Help() string {
	return "Usage: pwgen [ OPTIONS ] [ pw_length ] [ num_pw ]\n\nOptions supported by pwgen:\n  -c or --capitalize\n\tInclude at least one capital letter in the password\n  -A or --no-capitalize\n\tDon't include capital letters in the password\n  -n or --numerals\n\tInclude at least one number in the password\n  -0 or --no-numerals\n\tDon't include numbers in the password\n  -y or --symbols\n\tInclude at least one special symbol in the password\n  -s or --secure\n\tGenerate completely random passwords\n  -B or --ambiguous\n\tDon't include ambiguous characters in the password\n  -h or --help\n\tPrint a help message\n  -H or --sha1=path/to/file[#seed]\n\tUse sha1 hash of given file as a (not so) random generator\n  -C\n\tPrint the generated passwords in columns\n  -1\n\tDon't print the generated passwords in columns\n  -v or --no-vowels\n\tDo not use any vowels so as to avoid accidental nasty words\n"
}

func GenerateDefaultPwd(length string) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		digits
	intLen, err := strconv.Atoi(length)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, intLen)
	buf[0] = digits[rand.Intn(len(digits))]
	for i := 1; i < intLen; i++ {
		buf[i] = all[rand.Intn(len(all))]
	}
	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})
	str := string(buf)
	return str
}

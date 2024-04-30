package main

import (
	"flag"
	"fmt"
	"github.com/edwardBurroughes/edgrep/internal/edgrep"
	"os"
)

func main() {
	recursiveFlag := flag.Bool("r", false, "recursive flag")
	excludeFlag := flag.Bool("v", false, "exclude a word")
	caseSensitivityFlag := flag.Bool("i", false, "case sensitivity flag")
	flag.Parse()

	args := flag.Args()
	argLen := len(args)
	if argLen != 2 {
		fmt.Println("Error:")
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	regex, file := args[0], args[1]
	regexMeta := edgrep.RegexMetaObject{Pattern: regex, Exclude: *excludeFlag, Casesensitvity: *caseSensitivityFlag}

	if !*recursiveFlag && edgrep.IsDirectory(file) {
		fmt.Printf("Uh oh you provided a file with the recursive")
		os.Exit(1)
	}

	if *recursiveFlag && edgrep.IsDirectory(file) || file == "*" {
		err := edgrep.RecursiveGrep(file, regexMeta, recursiveFlag)
		// errors for both cases then
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		grepOption := edgrep.GrepFileOptions{regexMeta, file, *recursiveFlag, ""}
		edgrep.GrepFile(grepOption, nil, nil)
	}
}

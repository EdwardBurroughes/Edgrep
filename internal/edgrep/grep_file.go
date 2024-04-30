package edgrep

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

func readFile(filePath string) (*os.File, error) {
	readFile, err := os.Open(filePath)
	if err != nil {
		errMsg := fmt.Sprintf("Error: Unable to read %s", filePath)
		return nil, errors.New(errMsg)
	}
	return readFile, nil
}

func createScannerFromFile(file *os.File) *bufio.Scanner {
	fileScan := bufio.NewScanner(file)
	fileScan.Split(bufio.ScanLines)
	return fileScan
}

func formatRecursiveLine(recursiveFlag bool, rootName, filePath, line string) string {
	if recursiveFlag && rootName != "" {
		pathLs := strings.Split(filePath, fmt.Sprintf("%s/", rootName))
		return fmt.Sprintf("%s:%s", pathLs[1], line)
	} else {
		return line
	}
}

func GrepFile(options GrepFileOptions, output chan<- string, wg *sync.WaitGroup) {
	file, err := readFile(options.Filepath)
	if err != nil {
		fmt.Println(err)
	}
	fileScan := createScannerFromFile(file)
	defer file.Close()

	for fileScan.Scan() {
		line := fileScan.Text()
		if MatchesPattern(options.Regexmeta, line) {
			line = formatRecursiveLine(options.Recursiveflag, options.Rootname, options.Filepath, line)
			if output != nil {
				output <- line
			} else {
				fmt.Println(line)
			}
		}
	}
	if wg != nil {
		wg.Done()
	}
}

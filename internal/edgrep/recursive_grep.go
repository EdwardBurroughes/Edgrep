package edgrep

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sync"
)

// add intergration test
func RecursiveGrep(file string, regex RegexMetaObject, recursiveFlag *bool) error {
	var wg sync.WaitGroup
	output := make(chan string)
	defer close(output)

	go func() {
		for line := range output {
			fmt.Println(line)
		}
	}()

	err := filepath.Walk(file,
		func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			wg.Add(1)
			grepOption := GrepFileOptions{regex, filePath, *recursiveFlag, path.Base(file)}
			go GrepFile(grepOption, output, &wg)
			return nil
		})
	if err != nil {
		return fmt.Errorf("failed to walk path %s: %w", file, err)
	}
	wg.Wait()
	return nil
}

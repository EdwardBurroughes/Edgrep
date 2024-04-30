//go:build integration
// +build integration

package edgrep

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"runtime"
	"slices"
	"sync"
	"testing"
)

func getTestFilePath() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	exeDir := filepath.Dir(filename)
	return filepath.Join(exeDir, "testdata", "test.txt"), nil
}

func buildGrepIntTest(regexMata RegexMetaObject, filePath string) []string {
	var outputLines []string
	testChan := make(chan string)
	var wg sync.WaitGroup
	options := GrepFileOptions{regexMata, filePath, false, ""}
	wg.Add(1)
	go func() {
		defer close(testChan)
		GrepFile(options, testChan, &wg)
	}()
	for line := range testChan {
		outputLines = append(outputLines, line)
	}
	slices.Sort(outputLines)
	return outputLines
}

type GrepFileTestStruct struct {
	rg       RegexMetaObject
	expected []string
}

func TestGrepFile(t *testing.T) {
	filePath, err := getTestFilePath()
	if err != nil {
		panic(err)
	}
	t.Run("filepath matches line", func(t *testing.T) {
		testEach := []GrepFileTestStruct{
			GrepFileTestStruct{RegexMetaObject{"hel", false, false}, []string{"hello world"}},
			GrepFileTestStruct{RegexMetaObject{"hel", false, true}, []string{"Hello world", "hello world"}},
			GrepFileTestStruct{RegexMetaObject{"hel", true, true}, []string{"10", "sausage"}},
			GrepFileTestStruct{RegexMetaObject{"/d", true, true}, []string{"10"}},
		}

		for _, testStruct := range testEach {
			result := buildGrepIntTest(testStruct.rg, filePath)
			assert.Equal(t, testStruct.expected, result)
		}
	})
}

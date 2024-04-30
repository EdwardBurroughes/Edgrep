package edgrep

import (
	"fmt"
	"os"
	"regexp"
)

func BuildCaseIsensitiveRegexStr(pattern string, caseSensitivity bool) string {
	var i string
	if caseSensitivity {
		i = "(?i)"
	}
	return fmt.Sprintf(`%s%s`, i, pattern)
}

func MatchesPattern(regexMeta RegexMetaObject, line string) bool {
	if regexMeta.Pattern == "" {
		return true
	}
	pattern := BuildCaseIsensitiveRegexStr(regexMeta.Pattern, regexMeta.Casesensitvity)
	match := regexp.MustCompile(pattern).MatchString(line)
	return regexMeta.Exclude != match
}

func IsDirectory(path string) bool {
	// not handling the err - cuff me
	fileInfo, _ := os.Stat(path)
	return fileInfo.IsDir()
}

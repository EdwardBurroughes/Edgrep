//go:build unit
// +build unit

package edgrep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatRecursiveLine(t *testing.T) {
	t.Run("recursive flag supplied and rootname non-empty str", func(t *testing.T) {
		result := formatRecursiveLine(true, "root_path", "root_path/test.txt", "hello world")
		assert.Equal(t, "test.txt:hello world", result)
	})
	t.Run("recursiveFlag false", func(t *testing.T) {
		result := formatRecursiveLine(false, "root_path", "root_path/test.txt", "hello world")
		assert.Equal(t, "hello world", result)
	})
	t.Run("root name is an empty string", func(t *testing.T) {
		result := formatRecursiveLine(true, "", "root_path/test.txt", "hello world")
		assert.Equal(t, "hello world", result)
	})
}

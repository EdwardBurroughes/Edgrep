//go:build unit
// +build unit

package edgrep

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchesPattern(t *testing.T) {
	const mockLine string = "hello world"
	t.Run("pattern is an empty string", func(t *testing.T) {
		rObj := RegexMetaObject{
			Pattern:        "",
			Casesensitvity: false,
			Exclude:        true,
		}
		assert.True(t, MatchesPattern(rObj, mockLine))
	})
	t.Run("matches mock line", func(t *testing.T) {
		strPatterns := []string{"hello", "world", "hello world"}
		for _, pattern := range strPatterns {
			rObj := RegexMetaObject{
				Pattern:        pattern,
				Casesensitvity: false,
				Exclude:        false,
			}
			assert.True(t, MatchesPattern(rObj, mockLine))
		}
	})
	t.Run("none matching pattern", func(t *testing.T) {
		strPatterns := []string{"sausage", "kippers", "fish"}
		for _, pattern := range strPatterns {
			rObj := RegexMetaObject{
				Pattern:        pattern,
				Casesensitvity: false,
				Exclude:        false,
			}
			assert.False(t, MatchesPattern(rObj, mockLine))
		}
	})
	t.Run("case insensitive match", func(t *testing.T) {
		rObj := RegexMetaObject{
			Pattern:        "HELLO",
			Casesensitvity: true,
			Exclude:        false,
		}
		assert.True(t, MatchesPattern(rObj, mockLine))
	})

	t.Run("exclude match case", func(t *testing.T) {
		rObj := RegexMetaObject{
			Pattern:        "hello",
			Casesensitvity: false,
			Exclude:        true,
		}
		assert.False(t, MatchesPattern(rObj, mockLine))
	})
	t.Run("case sensitive and exclude", func(t *testing.T) {
		rObj := RegexMetaObject{
			Pattern:        "HELLO",
			Casesensitvity: true,
			Exclude:        true,
		}
		assert.False(t, MatchesPattern(rObj, mockLine))
	})
}

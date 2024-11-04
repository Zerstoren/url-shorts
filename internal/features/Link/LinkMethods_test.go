package featureLink

import (
	"testing"
)

func TestGetNewCode(t *testing.T) {
	t.Run("Code: 1", func(t *testing.T) {
		code := getNewCode(1)
		if code != "b" {
			t.Error("Expected b, got ", code)
		}
	})

	t.Run("Code: 27", func(t *testing.T) {
		code := getNewCode(27)
		if code != "b" {
			t.Error("Expected b, got ", code)
		}
	})

	t.Run("Code: 61", func(t *testing.T) {
		code := getNewCode(61)
		if code != "b" {
			t.Error("Expected b, got ", code)
		}
	})
}

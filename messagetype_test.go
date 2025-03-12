package messagetype

import (
	"errors"
	"testing"
)

func TestDetectMessageType(t *testing.T) {
	t.Run("No arguments", func(t *testing.T) {
		result := DetectMessageType()
		if result != Normal {
			t.Errorf("Expected Normal for no arguments, got %v", result)
		}
	})

	t.Run("Empty string", func(t *testing.T) {
		result := DetectMessageType("")
		if result != Normal {
			t.Errorf("Expected Normal for empty string, got %v", result)
		}
	})

	t.Run("Error objects", func(t *testing.T) {
		result := DetectMessageType(errors.New("some error"))
		if result != Error {
			t.Errorf("Expected Error for error object, got %v", result)
		}

		// Nil error should be treated as Normal
		result = DetectMessageType(error(nil))
		if result != Normal {
			t.Errorf("Expected Normal for nil error, got %v", result)
		}
	})

	t.Run("Error keywords", func(t *testing.T) {
		errorKeywords := []string{
			"This is an error message",
			"Operation failed",
			"exit status 1",
			"variable undeclared",
			"function undefined",
			"fatal exception",
		}

		for _, keyword := range errorKeywords {
			result := DetectMessageType(keyword)
			if result != Error {
				t.Errorf("Expected Error for keyword %q, got %v", keyword, result)
			}
		}
	})
}

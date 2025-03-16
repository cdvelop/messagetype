package messagetype

import (
	"github.com/cdvelop/tinystring"
)

// Type represents the classification of message types in the system.
// Available types:
// - Normal (0): Standard message without special formatting
// - Info (1): Informational message usually displayed with distinct styling
// - Error (2): Error message indicating failures or critical issues
// - Warning (3): Warning message indicating potential issues
// - OK (4): Success or confirmation message
type Type uint8

const (
	Normal  Type = iota // 0
	Info                // 1
	Error               // 2
	Warning             // 3
	OK                  // 4
)

// DetectMessageType detects the message type based on its content
// Accepts multiple arguments of any type, processing strings and errors
func DetectMessageType(args ...any) Type {
	// If there are no arguments, return Normal
	if len(args) == 0 {
		return Normal
	}

	// Process each argument
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			if v == "" {
				continue
			}

			return detectFromString(v)

		case error:
			if v != nil {
				// Errors are always considered Error type
				return Error
			}
		}
	}

	return Normal
}

// detectFromString analyzes a string to determine the message type
func detectFromString(content string) Type {
	lowerContent := tinystring.Convert(content).ToLower().String()

	// Detect errors
	if tinystring.Contains(lowerContent, "error") ||
		tinystring.Contains(lowerContent, "failed") ||
		tinystring.Contains(lowerContent, "exit status 1") ||
		tinystring.Contains(lowerContent, "undeclared") ||
		tinystring.Contains(lowerContent, "undefined") ||
		tinystring.Contains(lowerContent, "fatal") {
		return Error
	}

	// Detect warnings
	if tinystring.Contains(lowerContent, "warning") ||
		tinystring.Contains(lowerContent, "warn") {
		return Warning
	}

	// Detect informational messages
	if tinystring.Contains(lowerContent, "info") ||
		tinystring.Contains(lowerContent, " ...") ||
		tinystring.Contains(lowerContent, "starting") ||
		tinystring.Contains(lowerContent, "initializing") ||
		tinystring.Contains(lowerContent, "success") {
		return Info
	}

	return Normal
}

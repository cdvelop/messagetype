# Message Type Package
<!-- START_SECTION:BADGES_SECTION -->
<a href="docs/img/badges.svg"><img src="docs/img/badges.svg" alt="Project Badges" title="Generated by badges.sh from github.com/cdvelop/devscripts"></a>
<!-- END_SECTION:BADGES_SECTION -->

## Overview
The `messagetype` package provides functionality for classifying text messages into different types based on their content. This package is useful for categorizing log messages, user notifications, or any text content that needs to be displayed with appropriate styling or handling.

## Message Types

The package defines the following message types:

- `Normal`: Standard message with no specific classification
- `Info`: Informational message
- `Error`: Error message
- `Warning`: Warning message
- `Success`: Success/confirmation message

## Usage

### Import the package

```go
import "github.com/cdvelop/messagetype"
```

### Using the message type detection

```go
message := "Failed to connect to database"
msgType := messagetype.DetectMessageType(message)

// msgType will be messagetype.Error

if msgType == messagetype.Error {
    // Handle error message
}

// Multiple arguments
msgType := messagetype.DetectMessageType("Processing", err, "Completed")
// Returns Error (highest priority)
```

### Using helper methods for type checking

The package provides convenient helper methods for checking message types:

```go
message := "Starting application..."
msgType := messagetype.DetectMessageType(message)

// Using helper methods (recommended)
if msgType.IsInfo() {
    println("This is an informational message")
}

if msgType.IsError() {
    println("This is an error message")
}

if msgType.IsWarning() {
    println("This is a warning message")
}

if msgType.IsSuccess() {
    println("This is a success message")
}

if msgType.IsNormal() {
    println("This is a normal message")
}

// Get string representation
println("Message type:", msgType.String())
```

### Available helper methods

- `IsNormal()` - Returns true if the message type is Normal
- `IsInfo()` - Returns true if the message type is Info  
- `IsError()` - Returns true if the message type is Error
- `IsWarning()` - Returns true if the message type is Warning
- `IsSuccess()` - Returns true if the message type is Success
- `String()` - Returns a string representation of the message type

## Detection Logic

The `DetectMessageType` function analyzes the content of a message and categorizes it based on keywords:

- **Error messages**: Contain keywords like "error", "failed", "exit status 1", "undeclared", "undefined", or "fatal"
- **Warning messages**: Contain keywords like "warning" or "warn"
- **Info messages**: Contain keywords like "info", "...", "starting", "initializing", or "success"
- **Normal messages**: Default type for messages that don't match other categories

## Dependencies

This package depends on:
- [github.com/cdvelop/tinystring](https://github.com/cdvelop/tinystring) - For efficient string operations optimized for TinyGo/WebAssembly environments and minimal binary size

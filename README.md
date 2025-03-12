# MessageType Package

## Overview
The `messagetype` package provides functionality for classifying text messages into different types based on their content. This package is useful for categorizing log messages, user notifications, or any text content that needs to be displayed with appropriate styling or handling.

## Message Types

The package defines the following message types:

- `Normal`: Standard message with no specific classification
- `Info`: Informational message
- `Error`: Error message
- `Warning`: Warning message
- `OK`: Success/confirmation message

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
```

## Detection Logic

The `DetectMessageType` function analyzes the content of a message and categorizes it based on keywords:

- **Error messages**: Contain keywords like "error", "failed", "exit status 1", "undeclared", "undefined", or "fatal"
- **Warning messages**: Contain keywords like "warning" or "warn"
- **Info messages**: Contain keywords like "info", "...", "starting", "initializing", or "success"
- **Normal messages**: Default type for messages that don't match other categories

## Dependencies

This package depends on:
- [github.com/cdvelop/tinystring](https://github.com/cdvelop/tinystring) - For efficient string operations

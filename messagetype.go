package messagetype

import (
	"github.com/cdvelop/tinystring"
)

// messageType define el tipo de mensaje
type messageType string

const (
	Normal  messageType = "normal"
	Info    messageType = "info"
	Error   messageType = "error"
	Warning messageType = "warn"
	OK      messageType = "ok"
)

// Función para detectar el tipo de mensaje basado en su contenido
func DetectMessageType(content string) messageType {
	lowerContent := tinystring.Convert(content).ToLower().String()

	// Detectar errores
	if tinystring.Contains(lowerContent, "error") ||
		tinystring.Contains(lowerContent, "failed") ||
		tinystring.Contains(lowerContent, "exit status 1") ||
		tinystring.Contains(lowerContent, "undeclared") ||
		tinystring.Contains(lowerContent, "undefined") ||
		tinystring.Contains(lowerContent, "fatal") {
		return Error
	}

	// Detectar advertencias
	if tinystring.Contains(lowerContent, "warning") ||
		tinystring.Contains(lowerContent, "warn") {
		return Warning
	}

	// Detectar mensajes informativos
	if tinystring.Contains(lowerContent, "info") ||
		tinystring.Contains(lowerContent, " ...") ||
		tinystring.Contains(lowerContent, "starting") ||
		tinystring.Contains(lowerContent, "initializing") ||
		tinystring.Contains(lowerContent, "success") {
		return Info
	}

	return Normal
}

package messagetype

import (
	"github.com/cdvelop/tinystring"
)

// MessageType define el tipo de mensaje
type MessageType uint8

const (
	Normal  MessageType = iota // 0
	Info                       // 1
	Error                      // 2
	Warning                    // 3
	OK                         // 4
)

// DetectMessageType detecta el tipo de mensaje basado en su contenido
// Acepta múltiples argumentos de cualquier tipo, procesando strings y errores
func DetectMessageType(args ...any) MessageType {
	// Si no hay argumentos, retornar Normal
	if len(args) == 0 {
		return Normal
	}

	// Procesar cada argumento
	for _, arg := range args {
		switch v := arg.(type) {
		case string:
			if v == "" {
				continue
			}

			return detectFromString(v)

		case error:
			if v != nil {
				// Los errores siempre se consideran de tipo Error
				return Error
			}
		}
	}

	return Normal
}

// detectFromString analiza una cadena para determinar el tipo de mensaje
func detectFromString(content string) MessageType {
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

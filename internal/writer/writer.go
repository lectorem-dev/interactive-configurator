package writer

import (
	"path/filepath"
	"runtime"
	"strings"
)

// SaveValue — универсальная функция записи значения в файл.
// Автоматически определяет формат файла по расширению и вызывает соответствующий writer.
func SaveValue(filePath, key, value string) error {
	// Определение OS (можно использовать при необходимости)
	_ = runtime.GOOS

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".ini":
		return WriteINI(filePath, key, value)
	case ".json":
		return WriteJSON(filePath, key, value)
	default:
		return WritePlain(filePath, key, value)
	}
}

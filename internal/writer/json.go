package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// WriteJSON — обновляет ключ в JSON-файле
func WriteJSON(filePath, key, value string) error {
	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла %s: %w", filePath, err)
	}
	defer file.Close()

	var data map[string]interface{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	// Простое обновление ключа с поддержкой вложенности через "."
	parts := strings.Split(key, ".")
	last := len(parts) - 1
	m := data
	for i, part := range parts {
		if i == last {
			m[part] = value
		} else {
			if _, ok := m[part]; !ok {
				m[part] = map[string]interface{}{}
			}
			m = m[part].(map[string]interface{})
		}
	}

	// Перезапись файла
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("ошибка записи JSON: %w", err)
	}

	return nil
}

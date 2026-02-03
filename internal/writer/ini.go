package writer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// WriteINI — записывает или обновляет ключ=значение в ini-файле
func WriteINI(filePath, key, value string) error {
	// Проверка существования файла
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла %s: %w", filePath, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			line = fmt.Sprintf("%s=%s", key, value)
			found = true
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	if !found {
		lines = append(lines, fmt.Sprintf("%s=%s", key, value))
	}

	// Перезапись файла
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer outFile.Close()

	for _, line := range lines {
		if _, err := fmt.Fprintln(outFile, line); err != nil {
			return fmt.Errorf("ошибка записи в файл: %w", err)
		}
	}

	return nil
}

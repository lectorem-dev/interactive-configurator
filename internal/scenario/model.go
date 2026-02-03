package scenario

import (
	"errors"
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"strings"
)

// Scenario — весь сценарий установки
type Scenario struct {
	Steps []Step `json:"steps"`
}

// Step — один шаг сценария
type Step struct {
	File       string    `json:"file"`                 // путь к конфигу
	Key        string    `json:"key"`                  // имя переменной (ключ)
	Type       ValueType `json:"type"`                 // тип значения
	Comment    string    `json:"comment,omitempty"`    // подсказка оператору
	EnumValues []string  `json:"enumValues,omitempty"` // список допустимых значений для enum
	Default    *string   `json:"default,omitempty"`    // значение по умолчанию
}

// Validate проверяет корректность сценария
func (sc *Scenario) Validate() error {
	if len(sc.Steps) == 0 {
		return errors.New("сценарий не содержит шагов")
	}

	for i, step := range sc.Steps {
		if step.File == "" {
			return fmt.Errorf("шаг %d: путь к файлу не указан", i+1)
		}
		if step.Key == "" {
			return fmt.Errorf("шаг %d: ключ переменной не указан", i+1)
		}
		if !step.Type.IsValid() {
			return fmt.Errorf("шаг %d: тип значения '%s' не поддерживается", i+1, step.Type)
		}
		// для enum нужно хотя бы одно допустимое значение
		if step.Type == TypeEnum && len(step.EnumValues) == 0 {
			return fmt.Errorf("шаг %d: для типа enum должны быть указаны допустимые значения", i+1)
		}
		// проверка пути файла на корректность (необязательная)
		if _, err := filepath.Abs(step.File); err != nil {
			return fmt.Errorf("шаг %d: путь к файлу некорректен: %w", i+1, err)
		}
	}
	return nil
}

// IsValid проверяет, что тип значения поддерживается
func (t ValueType) IsValid() bool {
	switch t {
	case TypeString, TypeInt, TypeBool, TypeIP, TypePort, TypeAny, TypeEnum, TypePath:
		return true
	}
	return false
}

// Проверка конкретного значения по типу (опционально, для CLI)
func (t ValueType) ValidateValue(input string, enumValues []string) error {
	switch t {
	case TypeString, TypeAny:
		return nil
	case TypeInt:
		if _, err := strconv.Atoi(input); err != nil {
			return fmt.Errorf("ожидается целое число, получили '%s'", input)
		}
	case TypeBool:
		lower := strings.ToLower(input)
		if lower != "true" && lower != "false" {
			return fmt.Errorf("ожидается true или false, получили '%s'", input)
		}
	case TypeIP:
		if net.ParseIP(input) == nil {
			return fmt.Errorf("невалидный IP: %s", input)
		}
	case TypePort:
		p, err := strconv.Atoi(input)
		if err != nil || p < 1 || p > 65535 {
			return fmt.Errorf("порт должен быть 1-65535, получили '%s'", input)
		}
	case TypePath:
		if _, err := filepath.Abs(input); err != nil {
			return fmt.Errorf("невалидный путь: %s", input)
		}
	case TypeEnum:
		for _, val := range enumValues {
			if input == val {
				return nil
			}
		}
		return fmt.Errorf("значение должно быть одним из: %v", enumValues)
	default:
		return fmt.Errorf("неизвестный тип: %s", t)
	}
	return nil
}

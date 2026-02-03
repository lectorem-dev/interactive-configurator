package validate

import (
	"errors"
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"strings"
)

// GetValidator возвращает валидатор для указанного типа значения.
// valueType — тип значения: int, bool, ip, port, path, string/any
func GetValidator(valueType string) Validator {
	switch strings.ToLower(valueType) {
	case "int":
		return IntValidator{}
	case "bool":
		return BoolValidator{}
	case "ip":
		return IPValidator{}
	case "port":
		return PortValidator{}
	case "path":
		return PathValidator{}
	case "string", "any":
		return AnyValidator{}
	default:
		// По умолчанию не проверяем (любой ввод допустим)
		return AnyValidator{}
	}
}

// AnyValidator — валидатор для типа string/any (все значения допустимы)
type AnyValidator struct{}

// Validate всегда возвращает nil — любой ввод допустим
func (v AnyValidator) Validate(input string) error {
	return nil
}

// IntValidator — проверка, что введено целое число
type IntValidator struct{}

func (v IntValidator) Validate(input string) error {
	if _, err := strconv.Atoi(input); err != nil {
		return fmt.Errorf("неверное целое число: %s", input)
	}
	return nil
}

// BoolValidator — проверка, что введено true или false
type BoolValidator struct{}

func (v BoolValidator) Validate(input string) error {
	lower := strings.ToLower(input)
	if lower != "true" && lower != "false" {
		return errors.New("ожидается true или false")
	}
	return nil
}

// IPValidator — проверка корректности IP адреса (IPv4 или IPv6)
type IPValidator struct{}

func (v IPValidator) Validate(input string) error {
	if net.ParseIP(input) == nil {
		return fmt.Errorf("невалидный IP: %s", input)
	}
	return nil
}

// PortValidator — проверка диапазона порта 1-65535
type PortValidator struct{}

func (v PortValidator) Validate(input string) error {
	p, err := strconv.Atoi(input)
	if err != nil || p < 1 || p > 65535 {
		return fmt.Errorf("порт должен быть 1-65535, получили %s", input)
	}
	return nil
}

// PathValidator — проверка корректности пути к файлу/директории
type PathValidator struct{}

func (v PathValidator) Validate(input string) error {
	if _, err := filepath.Abs(input); err != nil {
		return fmt.Errorf("невалидный путь: %s", input)
	}
	return nil
}

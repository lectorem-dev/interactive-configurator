package validate

import (
	"errors"
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"strings"
)

func GetValidator(valueType string, enumValues []string) Validator {
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
	case "enum":
		return EnumValidator{Values: enumValues}
	case "string", "any":
		return AnyValidator{}
	default:
		return AnyValidator{}
	}
}

// Валидаторы

type AnyValidator struct{}

func (v AnyValidator) Validate(input string) error {
	return nil
}

type IntValidator struct{}

func (v IntValidator) Validate(input string) error {
	if _, err := strconv.Atoi(input); err != nil {
		return fmt.Errorf("неверное целое число: %s", input)
	}
	return nil
}

type BoolValidator struct{}

func (v BoolValidator) Validate(input string) error {
	lower := strings.ToLower(input)
	if lower != "true" && lower != "false" {
		return errors.New("ожидается true или false")
	}
	return nil
}

type IPValidator struct{}

func (v IPValidator) Validate(input string) error {
	if net.ParseIP(input) == nil {
		return fmt.Errorf("невалидный IP: %s", input)
	}
	return nil
}

type PortValidator struct{}

func (v PortValidator) Validate(input string) error {
	p, err := strconv.Atoi(input)
	if err != nil || p < 1 || p > 65535 {
		return fmt.Errorf("порт должен быть 1-65535, получили %s", input)
	}
	return nil
}

type PathValidator struct{}

func (v PathValidator) Validate(input string) error {
	if _, err := filepath.Abs(input); err != nil {
		return fmt.Errorf("невалидный путь: %s", input)
	}
	return nil
}

type EnumValidator struct {
	Values []string
}

func (v EnumValidator) Validate(input string) error {
	for _, val := range v.Values {
		if input == val {
			return nil
		}
	}
	return fmt.Errorf("значение должно быть одним из: %v", v.Values)
}

package validate

// Validator — интерфейс для проверки корректности введённого значения
type Validator interface {
	Validate(input string) error
}

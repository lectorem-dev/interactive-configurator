package validate

type Validator interface {
	Validate(input string) error
}

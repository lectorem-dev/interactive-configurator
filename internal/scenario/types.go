package scenario

type ValueType string

const (
	TypeString ValueType = "string" // любая строка
	TypeInt    ValueType = "int"    // целое число
	TypeBool   ValueType = "bool"   // true/false
	TypeIP     ValueType = "ip"     // IPv4 или IPv6
	TypePort   ValueType = "port"   // 1-65535
	TypeAny    ValueType = "any"    // без проверок
	TypeEnum   ValueType = "enum"   // ограниченный список значений
	TypePath   ValueType = "path"   // путь к файлу/директории
)

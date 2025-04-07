package utils

import (
	"regexp"
	"strings"
)

// ToTitleCase convierte una cadena a formato de título.
// ToTitleCase converts the input string `s` to title case, where the first letter
// of each word is capitalized and the rest are in lowercase.
// It is useful for formatting strings to improve readability and consistency.
//
// Parameters:
//   - s: The input string to be converted to title case.
//
// Returns:
//   - A new string where each word starts with an uppercase letter and the remaining
//     letters are in lowercase.
func ToTitleCase(s string) string {
	return strings.Title(s)
}

// IsEmpty verifica si una cadena está vacía.
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// Contains verifica si una cadena contiene un substring.
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// IsValidEmail verifica si un correo electrónico tiene un formato válido.
// Devuelve true si el correo es válido, de lo contrario false.
func IsValidEmail(email string) bool {
	// Expresión regular para validar correos electrónicos
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

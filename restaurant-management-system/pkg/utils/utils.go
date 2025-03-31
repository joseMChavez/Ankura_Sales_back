package utils

import "strings"

// ToTitleCase convierte una cadena a formato de título.
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
package common

import (
	"strings"
	"unicode"
)

func LowercaseAndRemovePunctuation(input string) string {
	// Mengubah semua karakter menjadi lowercase
	lowered := strings.ToLower(input)

	// Membuat builder untuk menyusun hasil string tanpa tanda baca
	var result strings.Builder

	// Loop setiap karakter dalam string
	for _, char := range lowered {
		// Menambahkan karakter ke hasil jika bukan tanda baca
		if !unicode.IsPunct(char) {
			result.WriteRune(char)
		}
	}

	// Mengembalikan hasil dalam bentuk string
	return result.String()
}

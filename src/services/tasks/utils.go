package tasks

import "unicode"

// CountCharTypes
// Count the number of vowels, consonants and other types of chars.
func CountCharTypes(message string) (vowels int, consonants int, nonLetters int) {
	for _, c := range message {
		if !unicode.IsLetter(c) {
			nonLetters++
		} else if IsVowel(c) {
			vowels++
		} else {
			consonants++
		}
	}

	return
}

// IsVowel
// Test if a character is a vowel or not
func IsVowel(character rune) bool {
	switch unicode.ToLower(character) {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}

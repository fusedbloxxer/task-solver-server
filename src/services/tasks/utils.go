package tasks

import (
	"sort"
	"unicode"
)

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

func IsPalindrome(message string) bool {
	for i := 0; i != len(message)/2; i++ {
		if message[i] != message[len(message)-i-1] {
			return false
		}
	}

	return true
}

// IsBirdLanguage
// Tests if a string ends with the letter "p".
func IsBirdLanguage(message string) bool {
	for i, c := range message {
		if i == len(message)-1 {
			return !IsVowel(c)
		}

		if IsVowel(c) && message[i+1] != 'p' {
			return false
		}
	}

	return true
}

func IsVowelSymmetric(message string) bool {
	if len(message) == 0 {
		return false
	}

	return IsVowel(rune(message[0])) && IsVowel(rune(message[len(message)-1]))
}

func IsAnagramTo(message string, anagram string) bool {
	rawComparisonMessage := []byte(anagram)
	rawInputMessage := []byte(message)

	sort.Slice(rawInputMessage, func(i, j int) bool {
		return rawInputMessage[i] < rawInputMessage[j]
	})

	sort.Slice(rawComparisonMessage, func(i, j int) bool {
		return rawComparisonMessage[i] < rawComparisonMessage[j]
	})

	return string(rawComparisonMessage) == string(rawInputMessage)
}

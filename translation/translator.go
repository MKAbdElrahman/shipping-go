package translation

import "strings"

// Translate translates a word into a language
// Supported words: hello
// Supported languages: english, german
func Translate(word string, language string) string {

	word = sanitizeInput(word)
	language = sanitizeInput(language)

	if word != "hello" {
		return "unsupported word"
	}
	switch language {
	case "english":
		return "hello"
	case "german":
		return "hallo"
	default:
		return "unsupported language"
	}
}

func sanitizeInput(word string) string {
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	return word
}

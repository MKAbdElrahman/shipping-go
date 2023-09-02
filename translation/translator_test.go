package translation_test

import (
	"shipping-go/translation"
	"testing"
)

func TestTranslate(t *testing.T) {
	// Arrange

	tt := []struct {
		word     string
		language string
		expected string
	}{
		{"hello", "english", "hello"},
		{"hello", "german", "hallo"},
		{"hello", "french", "unsupported language"},
		{" Hello", "English  ", "hello"},
	}
	// Act
	for _, tc := range tt {
		actual := translation.Translate(tc.word, tc.language)
		// Assert
		if actual != tc.expected {
			t.Errorf("Translate(%s, %s): expected %s, actual %s",
				tc.word, tc.language, tc.expected, actual)
		}
	}

}

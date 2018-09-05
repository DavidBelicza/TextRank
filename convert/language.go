package convert

import "unicode/utf8"

// Language interface and its methods make possible the polimorf usage of
// language specific features by custom implementations.
type Language interface {
	IsStopWord(word string) bool
	FindRootWord(word string) (bool, string)
	SetActiveLanguage(code string)
	SetWords(code string, words []string)
}

// LanguageDefault struct is implementation of Language interface. It stores
// the stop words of loaded languages and can find stop words by tokens.
type LanguageDefault struct {
	defaultLang string
	languages   map[string][]string
}

// NewLanguage constructor of the LanguageDefault Retrieves a pointer
// LanguageDefault. It has setup to English by default.
func NewLanguage() *LanguageDefault {
	lang := &LanguageDefault{
		"en",
		make(map[string][]string),
	}

	words := getDefaultEnglish()

	lang.SetWords("en", words)

	return lang
}

// IsStopWord method retrieves true when the given word is in the stop word
// list or when the word has less character then 2.
func (lang *LanguageDefault) IsStopWord(word string) bool {
	if utf8.RuneCountInString(word) <= 2 {
		return true
	}

	if stopWords, ok := lang.languages[lang.defaultLang]; ok {
		for _, val := range stopWords {
			if val == word {
				return true
			}
		}
	}

	return false
}

// FindRootWord method gets a word as an input, "apples" for example and it
// retrieves the root-word of this given word, "apple" for example. The first
// return parameter is true when a word-root has found, otherwise it's false.
func (lang *LanguageDefault) FindRootWord(word string) (bool, string) {
	return false, ""
}

// SetActiveLanguage method switch between languages by the language's code. The
// language code is not standard, it can be anything.
func (lang *LanguageDefault) SetActiveLanguage(code string) {
	lang.defaultLang = code
}

// SetWords method set stop words into the LanguageDefault struct by the
// language's code.
func (lang *LanguageDefault) SetWords(code string, words []string) {
	lang.languages[code] = words
}

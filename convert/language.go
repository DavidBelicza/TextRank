package convert

type Language struct {
	defaultLang string
	languages   map[string][]string
}

func NewLanguage() *Language {
	lang := &Language{
		"en",
		make(map[string][]string),
	}

	words := getDefaultEnglish()

	lang.SetWords("en", words)

	return lang
}

func (lang *Language) IsStopWord(word string) bool {
	if stopWords, ok := lang.languages[lang.defaultLang]; ok {
		for _, val := range stopWords {
			if val == word {
				return true
			}
		}
	}

	return false
}

func (lang *Language) SetDefaultLanguage(code string) {
	lang.defaultLang = code
}

func (lang *Language) SetWords(code string, words []string) {
	lang.languages[code] = words
}

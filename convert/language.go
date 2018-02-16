package convert

type Language interface {
	IsStopWord(word string) bool
	SetActiveLanguage(code string)
	SetWords(code string, words []string)
}

type LanguageDefault struct {
	defaultLang string
	languages   map[string][]string
}

func NewLanguage() *LanguageDefault {
	lang := &LanguageDefault{
		"en",
		make(map[string][]string),
	}

	words := getDefaultEnglish()

	lang.SetWords("en", words)

	return lang
}

func (lang *LanguageDefault) IsStopWord(word string) bool {
	if len(word) <= 2 {
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

func (lang *LanguageDefault) SetActiveLanguage(code string) {
	lang.defaultLang = code
}

func (lang *LanguageDefault) SetWords(code string, words []string) {
	lang.languages[code] = words
}

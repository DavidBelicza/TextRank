package parse

type Rule interface {
	IsWordSeparator(rune rune) bool
	IsSentenceSeparator(rune rune) bool
}

type RuleDefault struct {
	wordSeparators     [18]string
	sentenceSeparators [3]string
}

func NewRule() *RuleDefault {
	return &RuleDefault{
		[18]string{" ", ",", ")", "(", "[", "]", "{", "}", "\"", ";", "\n", ">", "<", "%", "@", "&", "=", "#"},
		[3]string{"!", ".", "?"},
	}
}

func (r *RuleDefault) IsWordSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range r.wordSeparators {
		if chr == val {
			return true
		}
	}

	return r.IsSentenceSeparator(rune)
}

func (r *RuleDefault) IsSentenceSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range r.sentenceSeparators {
		if chr == val {
			return true
		}
	}

	return false
}

package parse

// Rule interface and its methods make possible the polimorf usage of process
// how Rule retrieve tokens from text.
type Rule interface {
	IsWordSeparator(rune rune) bool
	IsSentenceSeparator(rune rune) bool
}

// RuleDefault struct implements the Rule interface. It contains the separator
// characters and can decide a character is separator or not.
type RuleDefault struct {
	wordSeparators     [21]string
	sentenceSeparators [3]string
}

// NewRule constructor retrieves a RuleDefault pointer.
func NewRule() *RuleDefault {
	return &RuleDefault{
		[21]string{" ", ",", "'", "’", "\"", ")", "(", "[", "]", "{", "}", "\"", ";", "\n", ">", "<", "%", "@", "&", "=", "#"},
		[3]string{"!", ".", "?"},
	}
}

// IsWordSeparator method retrieves true when a character is a kind of special
// character and possibly it separates to words from each other. It also checks
// for sentence separator by IsSentenceSeparator method.
func (r *RuleDefault) IsWordSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range r.wordSeparators {
		if chr == val {
			return true
		}
	}

	return r.IsSentenceSeparator(rune)
}

// IsSentenceSeparator method retrieves true when a character is a kind of
// special character and possibly it separates to words from each other.
func (r *RuleDefault) IsSentenceSeparator(rune rune) bool {
	chr := string(rune)

	for _, val := range r.sentenceSeparators {
		if chr == val {
			return true
		}
	}

	return false
}

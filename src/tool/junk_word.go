package tool

//@todo list of junk words
func IsJunkWord(word string) bool {
	if word == "is" ||
		word == "a" ||
		word == "an" ||
		word == "in" ||
		word == "has" ||
		word == "the" ||
		word == "There" ||
		word == "It" {

		return true
	}

	return false
}

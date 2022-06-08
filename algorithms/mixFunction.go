package algorithms

func TestBool() int {
	variable := -5
	result := func(checker bool, acceptVal int, ignoreVal int) int {
		if checker {
			return acceptVal
		}
		return ignoreVal
	}(variable > 0, variable, -variable)
	return result
}

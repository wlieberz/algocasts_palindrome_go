package palindrome

func Palindrome(input string) bool {
	var frontIdx int
	var backIdx int
	j := len(input)
	i := 1

	for i != j {
		frontIdx = i - 1
		backIdx = j - 1
		frontChar := string([]rune(input)[frontIdx])
		backChar := string([]rune(input)[backIdx])

		if frontChar != backChar {
			return false
		}

		// Prevent infinite loop: determine when we've compared all values:
		if i+1 == j {
			break
		}

		i += 1
		j -= 1
	}

	return true
}

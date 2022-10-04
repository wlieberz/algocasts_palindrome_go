package palindrome

func Palindrome(input string) bool {
	var frontIdx int
	var backIdx int
	j := len(input)
	i := 1

	for rangeIdx, _ := range input {
		frontIdx = i - 1
		backIdx = j - 1
		frontChar := string([]rune(input)[frontIdx])
		backChar := string([]rune(input)[backIdx])

		if frontChar != backChar {
			return false
		}

		// If i is "one away" from j, we've compared all values:
		if i+1 == j {
			break
		}

		i += 1
		j -= 1
		// "consume" rangeIdx via discard to make the Go compiler happy:
		_ = rangeIdx
	}

	return true
}

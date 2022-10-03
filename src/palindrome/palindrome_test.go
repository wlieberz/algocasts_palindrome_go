package palindrome

import "testing"

func Test_aba_is_palindrome(t *testing.T) {
	input := "aba"
	expected := true

	output := Palindrome(input)
	if output != expected {
		t.Fatalf("EXPECTED: %v GOT: %v", expected, output)
	}
}

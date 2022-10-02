package palindrome

import "testing"

func aba_is_palindrome(t *testing.T) {
	input := "aba"
	expected := true
	output := Palindrome(input)
	if output != expected {
		t.Fatalf("EXPECTED: %q\n  GOT: %q\n", expected, output)
	}
}

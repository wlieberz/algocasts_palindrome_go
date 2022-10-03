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

func Test_aba_space_before_not_palindrome(t *testing.T) {
	input := " aba"
	expected := false

	output := Palindrome(input)
	if output != expected {
		t.Fatalf("EXPECTED: %v GOT: %v", expected, output)
	}
}

func Test_aba_space_after_not_palindrome(t *testing.T) {
	input := "aba "
	expected := false

	output := Palindrome(input)
	if output != expected {
		t.Fatalf("EXPECTED: %v GOT: %v", expected, output)
	}
}

func Test_greetings_not_palindrome(t *testing.T) {
	input := "greetings"
	expected := false

	output := Palindrome(input)
	if output != expected {
		t.Fatalf("EXPECTED: %v GOT: %v", expected, output)
	}
}

func Test_numeric_is_palindrome(t *testing.T) {
	input := "1000000001"
	expected := true

	output := Palindrome(input)
	if output != expected {
		t.Fatalf("EXPECTED: %v GOT: %v", expected, output)
	}
}

func Test_fish_different_casing_not_palindrome(t *testing.T) {
	input := "Fish hsif"
	expected := false

	output := Palindrome(input)
	if output != expected {
		t.Fatalf("EXPECTED: %v GOT: %v", expected, output)
	}
}

func Test_pennep_is_palindrome(t *testing.T) {
	input := "pennep"
	expected := true

	output := Palindrome(input)
	if output != expected {
		t.Fatalf("EXPECTED: %v GOT: %v", expected, output)
	}
}

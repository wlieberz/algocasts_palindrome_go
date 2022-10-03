# About

This repo contains my solution for a palindrome programming exercise.

There are many variations of this type of exercise, with varying levels of
complexity.

This particular variation comes from the Udemy course [The Coding Inteview Bootcamp: Algorithms + Data Structures](https://www.udemy.com/course/coding-interview-bootcamp-algorithms-and-data-structure/).

The companion repo for the course can be found here:
[StephenGrider/AlgoCasts](https://github.com/StephenGrider/AlgoCasts)

This solution is in Go.

# Problem Description

>--- Directions:
>
>Given a string, return true if the string is a palindrome
>or false if it is not.  Palindromes are strings that
>form the same word if it is reversed. 
>
> *Do* include spaces and punctuation in determining if the string is a palindrome.
>
>--- Examples:
>
>   palindrome("abba") === true
>
>   palindrome("abcdefg") === false

# Running Tests

Prerequisite: you should have the [Go installed](https://go.dev/doc/install).

In order to run the **unit-tests**:

```sh
cd src/palindrome
go test -v
```
**Expected output:**
```sh
=== RUN   Test_aba_is_palindrome
--- PASS: Test_aba_is_palindrome (0.00s)
=== RUN   Test_aba_space_before_not_palindrome
--- PASS: Test_aba_space_before_not_palindrome (0.00s)
=== RUN   Test_aba_space_after_not_palindrome
--- PASS: Test_aba_space_after_not_palindrome (0.00s)
=== RUN   Test_greetings_not_palindrome
--- PASS: Test_greetings_not_palindrome (0.00s)
=== RUN   Test_numeric_is_palindrome
--- PASS: Test_numeric_is_palindrome (0.00s)
=== RUN   Test_fish_different_casing_not_palindrome
--- PASS: Test_fish_different_casing_not_palindrome (0.00s)
=== RUN   Test_pennep_is_palindrome
--- PASS: Test_pennep_is_palindrome (0.00s)
PASS
```
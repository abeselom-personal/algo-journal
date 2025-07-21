package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abad"))

}

func longestPalindrome(s string) string {
	// first condition check
	if len(s) == 0 {
		return ""
	}
	expand := func(left, right int) string {
		for left >= 0 && right < len(s) && s[left] == s[right] {

			left--
			right++
		}
		return s[left+1 : right]
	}

	longest := ""
	for i := range s {
		odd := expand(i, i)
		even := expand(i, i+1)
		if len(even) > len(longest) {
			longest = odd
		}
		if len(odd) > len(longest) {
			longest = odd
		}

	}
	return longest
}

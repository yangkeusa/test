// a package for palindromes

package palindrome

func IsPalinDrome(x string) bool {
	// stupid algorithm: store all the strings in an array and then check
	buffer := make([]int32, len(x))
	j := 0
	for _, rune := range x {
		buffer[j] = rune
		j++
	}
	a := 0
	z := j - 1
	for ; a < z; {
		if buffer[a] != buffer[z] {
			return false
		}
		a++
		z--
	}
	return true
}

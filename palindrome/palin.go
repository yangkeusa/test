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
	for a, z := 0, j-1 ; a < z; a, z = a+1, z-1 {
		if buffer[a] != buffer[z] {
			return false
		}
	}
	return true
}

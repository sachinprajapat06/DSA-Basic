package codes

func Palindrome(a string) bool {
	strlen := len(a) - 1
	for index, _ := range a {
		if index == strlen || index > strlen {
			return true
		}
		if a[index] != a[strlen] {
			return false
		}
		strlen--
	}
	return true
}

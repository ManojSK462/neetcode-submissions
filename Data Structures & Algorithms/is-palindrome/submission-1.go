
func isPalindrome(s string) bool {


	l :=0
	r := len(s)-1

	for l<r{
		for l<r && !isAlphaNum(s[l]){
			l++
		}
		for l<r && !isAlphaNum(s[r]){
			r--
		}
		if toLower(s[l]) != toLower(s[r]){
			return false
		}
		l++
		r--
		
	}
	return true
}
func isAlphaNum(c byte) bool {
    return (c >= 'a' && c <= 'z') ||
        (c >= 'A' && c <= 'Z') ||
        (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return c + 32
    }
    return c
}

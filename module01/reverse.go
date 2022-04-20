package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	// Strings-package solution - Passes all
	// s := strings.Split(word, "")
	// for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	// 	s[i], s[j] = s[j], s[i]
	// }
	// return strings.Join(s, "")

	// var sb strings.Builder - Fails Chinese characters
	// for i := len(word) - 1; i >= 0; i-- {
	// 	sb.WriteByte(word[i])
	// }
	// return sb.String()

	// Byte-base solution - Fails Chinese characters
	// var res string
	// for i := len(word) - 1; i >= 0; i-- {
	// 	res = res + string(word[i])
	// }
	// return res

	// Rune-based solution - Passes all
	var res string
	for _, r := range word {
		res = string(r) + res
	}
	return res
}

package module01

import (
// "fmt"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	if dec == 0 || base > 16 {
		return "0"
	}
	// More idiomatic implementation
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec / base
	}
	return res
}

// Recusion based solution
// func decToBaseOp(dec int, base int, res string) string {
// 	if dec == 0 {
// 		return res
// 	}
// 	rem := dec % base
// 	res = toBase(rem) + res
// 	num := dec / base
// 	return decToBaseOp(num, base, res)
// }
//
// func toBase(val int) string {
// 	hexMap := map[int]string{
// 		10: "A",
// 		11: "B",
// 		12: "C",
// 		13: "D",
// 		14: "E",
// 		15: "F",
// 	}
// 	if val < 10 {
// 		return fmt.Sprint(val)
// 	} else {
// 		return hexMap[val]
// 	}
// }

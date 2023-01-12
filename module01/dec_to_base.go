package module01

import (
	"strings"
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
const hex = "0123456789ABCDEF"

func DecToBase(dec, base int) string {
	var sb strings.Builder
	for dec > 0 {
		sb.WriteByte(hex[dec%base])
		dec = dec/base
	}
	return Reverse(sb.String())
}

func DecToBaseWithnamedReturn(dec, base int) (result string) {
	var sb strings.Builder
	for dec > 0 {
		sb.WriteByte(hex[dec%base])
		dec = dec/base
	}
	result = Reverse(sb.String())
	return 
}

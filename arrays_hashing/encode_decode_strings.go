package arrays_hashing

import (
	"strconv"
	"strings"
)

// Solution for Encode and Decode Strings
// Design an algorithm to encode a list of strings to a string. 
// The encoded string is then sent over the network and is decoded back to the original list of strings.

func Encode(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteByte('#')
		sb.WriteString(s)
	}
	return sb.String()
}

func Decode(s string) []string {
	var res []string
	i := 0
	for i < len(s) {
		j := i
		for s[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(s[i:j])
		res = append(res, s[j+1 : j+1+length])
		i = j + 1 + length
	}
	return res
}

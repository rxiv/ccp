package ccp

import (
	b64 "encoding/base64"
	"encoding/hex"
	"strings"

)

func ToHex(in string) string {
	return string(hex.EncodeToString([]byte(in)))
}

func FromHex(in string) string {
	out, _ := hex.DecodeString(in)
	return string(out)
}

func ToB64(in string) string {
	return b64.StdEncoding.EncodeToString([]byte(in))
}

func FromB64(in string) string {
	out, _ := b64.StdEncoding.DecodeString(in)
	return string(out)
}

func XorString(s1, key string) (out string) {
	for i := 0; i < len(s1); i++ {
		out += string(s1[i] ^ key[i%len(key)])
	}
	return out
}

func ScoreText(in string) (out int) {
	chars := "etaoin shrdlucmfwy"
	m := make(map[string]int)

	for i := 0; i < len(chars); i++ {
		val := strings.ToLower(string(chars[i]))
		m[val] = strings.Count(in, val)
	}
	var sum int
	for _, x := range m {
		sum += x
	}
	return sum
}

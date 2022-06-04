package github.com/rxiv/ccp

import (
	b64 "encoding/base64"
	"encoding/hex"
)

func to_hex(in string) string {
	return string(hex.EncodeToString([]byte(in)))
}

func from_hex(in string) string {
	out, _ := hex.DecodeString(in)
	return string(out)
}

func to_b64(in string) string {
	return b64.StdEncoding.EncodeToString([]byte(in))
}

func from_b64(in string) string {
	out, _ := b64.StdEncoding.DecodeString(in)
	return string(out)
}

func XorString(s1, key string) (out string) {
	for i := 0; i < len(s1); i++ {
		out += string(s1[i] ^ key[i%len(key)])
	}
	return out
}


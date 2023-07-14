package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func String2Md5(s string) (md5str string) {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func StringReplace(s string, rep string, new string, n int) string {
	return strings.Replace(s, rep, new, n)
}

func RandomString(length int) string {
	var chars = []byte("0123456789")
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		panic("Wrong charset length for NewLenChars()")
	}
	maxrb := 255 - (256 % clen)
	b := make([]byte, length)
	r := make([]byte, length+(length/4)) // storage for random bytes.
	i := 0
	for {
		if _, err := rand.Read(r); err != nil {
			panic("Error reading random bytes: " + err.Error())
		}
		for _, rb := range r {
			c := int(rb)
			if c > maxrb {
				continue // Skip c number to avoid modulo bias.
			}
			b[i] = chars[c%clen]
			i++
			if i == length {
				return string(b)
			}
		}
	}
}

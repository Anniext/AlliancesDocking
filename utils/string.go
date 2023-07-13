package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func String2Md5(s string) (md5str string) {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

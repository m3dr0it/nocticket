package util

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(stringToHash string) string {
	hash := md5.Sum([]byte(stringToHash))
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

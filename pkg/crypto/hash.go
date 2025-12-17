package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 計算字串的 MD5 雜湊值
func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// SHA256 計算字串的 SHA256 雜湊值
func SHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

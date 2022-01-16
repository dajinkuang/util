package cryptoutil

import (
	"crypto/hmac"
	"crypto/sha256"
)

// HmacSha256 -
func HmacSha256(key, data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

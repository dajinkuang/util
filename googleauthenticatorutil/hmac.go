package googleauthenticatorutil

import (
	"crypto/hmac"
	"crypto/sha1"
)

// HmacSha1 -
func HmacSha1(key, data []byte) []byte {
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}
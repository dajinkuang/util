package charactersetutil

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// UTF8ToGBK UTF8 转为 GBK
func UTF8ToGBK(b []byte) ([]byte, error) {
	result, _, err := transform.Bytes(simplifiedchinese.GB18030.NewEncoder(), b)
	return result, err
}

// GBKToUTF8 GBK 转为 UTF8
func GBKToUTF8(b []byte) ([]byte, error) {
	result, _, err := transform.Bytes(simplifiedchinese.GB18030.NewDecoder(), b)
	return result, err
}

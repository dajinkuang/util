package charactersetutil

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func UTF8ToGBK(b []byte) ([]byte, error) {
	result, _, err := transform.Bytes(simplifiedchinese.GB18030.NewEncoder(), b)
	return result, err
}

func GBKToUTF8(b []byte) ([]byte, error) {
	result, _, err := transform.Bytes(simplifiedchinese.GB18030.NewDecoder(), b)
	return result, err
}

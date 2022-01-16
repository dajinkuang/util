package phoneutil

import (
	"regexp"
	"strings"
)

// CheckPhone 检查手机号是否合法
func CheckPhone(phone string, countryCode string) bool {
	if v := FetchCountryByNumber(countryCode); v == nil {
		return false
	}
	pattern := "^[0-9]{5,}$"
	if strings.ToUpper(countryCode) == "+86" {
		pattern = "^[0-9]{11}$"
	}
	if b, err := regexp.Match(pattern, []byte(phone)); err != nil {
		return false
	} else if !b {
		return false
	}
	return true
}

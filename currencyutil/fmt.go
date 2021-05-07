package currencyutil

import "strings"

func FmtCurrencyNum(s string) string {
	arr := strings.Split(s, ".")
	if len(arr) < 2 {
		return arr[0]
	}
	if len(arr[1]) < 0 {
		return arr[0]
	}
	if arr[1] == "" {
		return arr[0]
	}
	var (
		ss      string
		runeArr = []rune(arr[1])
	)
	for i := len(runeArr) - 1; i > 0; i-- {
		if runeArr[i] == '0' {
			continue
		}
		ss = string(runeArr[0 : i+1])
		break
	}
	if ss == "" {
		return arr[0]
	} else {
		ss = arr[0] + "." + ss
	}

	return ss
}

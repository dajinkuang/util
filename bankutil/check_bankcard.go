package bankutil

// CheckBankcard 检查银行卡号是否合法
func CheckBankcard(bankcardCardNum string) bool {
	if len([]rune(bankcardCardNum)) < 9 {
		return false
	}
	for _, v := range []rune(bankcardCardNum) {
		if int64(v) < 48 || int64(v) > 57 {
			return false
		}
	}
	return true
}

package maskutil

// MaskPhone 隐去手机号中间 4 位地区码, 如 155****8818
func MaskPhone(phone string) string {
	if n := len(phone); n >= 8 {
		return phone[:n-8] + "****" + phone[n-4:]
	}
	return phone
}

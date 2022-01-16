package maskutil

// MaskIDCard 脱敏身份证号
// - 中国大陆地区 身份证 隐去身份证号中间 12 位, 如 1411************36
// - 台湾地区 身份证 台湾身份证总共有10位。第一位是字母。后面九位是数字。隐去身份证号中间 6 位
// - 其他国家及地区 护照 没有特定规律暂不处理
func MaskIDCard(idCard string) string {
	switch len(idCard) {
	case 10:
		return idCard[:2] + "******" + idCard[8:]
	case 18:
		return idCard[:4] + "************" + idCard[16:]
	}
	return idCard
}

// Package googleauthenticatorutil google 验证码工具包
package googleauthenticatorutil

// GenGoogleAuthCode 生成google验证码
func GenGoogleAuthCode(auth *GAuth, secret string) (code string, err error) {
	if auth == nil {
		auth = NewGAuth() // 默认长度是6
	}
	//auth.SetCodeLength(8) // 设置code长度
	code, err = auth.GetCode(secret)
	return
}

// VerifyGoogleCode 验证google验证码
func VerifyGoogleCode(auth *GAuth, secret, code string) (ok bool, err error) {
	if auth == nil {
		auth = NewGAuth() // 默认长度是6
	}
	// 1:30sec
	ok, err = auth.VerifyCode(secret, code, 1)
	if err != nil {
		return
	}
	return
}

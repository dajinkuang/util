package googleauthenticatorutil

func GenGoogleAuthCode(auth *GAuth, secret string) (code string, err error) {
	if auth == nil {
		auth = NewGAuth() // 默认长度是6
	}
	//auth.SetCodeLength(8) // 设置code长度
	code, err = auth.GetCode(secret)
	return
}

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

package mark2

func NewUserInfo() *UserInfo {
	ui := new(UserInfo)
	return ui
}

func NewResult() *Result {
	r := new(Result)
	r.Code = ResultCodes_NG
	r.Message = ""
	return r
}

func NewAccessToken() *AccessToken {
	at := new(AccessToken)
	return at
}

func NewLoginResult() *LoginResult {
	lr := new(LoginResult)
	lr.Result = NewResult()
	lr.AccessToken = NewAccessToken()
	return lr
}

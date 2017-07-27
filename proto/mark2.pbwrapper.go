package mark2

// ===== Model =====
func NewAccessToken() *AccessToken {
	at := new(AccessToken)
	return at
}

func NewUserInfo() *UserInfo {
	ui := new(UserInfo)
	return ui
}

func NewUserInfoList() *UserInfoList {
	uil := new(UserInfoList)
	return uil
}

func NewRoomInfo() *RoomInfo {
	ri := new(RoomInfo)
	return ri
}

func NewRoomInfoList() *RoomInfoList {
	ril := new(RoomInfoList)
	return ril
}

func NewMessage() *Message {
	m := new(Message)
	return m
}

// ===== Request =====

// ===== Response =====
func NewResult() *Result {
	r := new(Result)
	r.Code = ResultCodes_NG
	r.Message = ""
	return r
}

func NewLoginResult() *LoginResult {
	lr := new(LoginResult)
	lr.Result = NewResult()
	lr.AccessToken = NewAccessToken()
	return lr
}

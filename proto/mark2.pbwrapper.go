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

func NewRoomInfo() *RoomInfo {
	ri := new(RoomInfo)
	return ri
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
	r := new(LoginResult)
	r.Result = NewResult()
	r.AccessToken = NewAccessToken()
	return r
}

func NewUserInfoResult() *UserInfoResult {
	r := new(UserInfoResult)
	r.Result = NewResult()
	r.UserInfoList = make([]*UserInfo, 0)
	return r
}

func NewRoomInfoResult() *RoomInfoResult {
	r := new(RoomInfoResult)
	r.Result = NewResult()
	r.RoomInfoList = make([]*RoomInfo, 0)
	return r
}

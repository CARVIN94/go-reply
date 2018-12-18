package reply

// Urlencoded 错误
func Urlencoded() *Error {
	return &Error{Code: 1001, Message: "is not Urlencoded"}
}

// Multipart 错误
func Multipart() *Error {
	return &Error{Code: 1002, Message: "is not Multipart"}
}

// SignInError 错误
func SignInError() *Error {
	return &Error{Code: 1003, Message: "账号或密码错误"}
}

// TokenEncryptError 错误
func TokenEncryptError() *Error {
	return &Error{Code: 1004, Message: "加密错误"}
}

// TokenDecryptError 错误
func TokenDecryptError() *Error {
	return &Error{Code: 1005, Message: "密钥错误"}
}

// AuthError 错误
func AuthError() *Error {
	return &Error{Code: 1006, Message: "权限不足"}
}

// NotExist 错误
func NotExist(name string) *Error {
	return &Error{Code: 2001, Message: "参数不存在: " + name}
}

// NotInt 错误
func NotInt(name string) *Error {
	return &Error{Code: 2002, Message: "参数类型错误（int）: " + name}
}

// NotExistByAccount 账号不存在
func NotExistByAccount() *Error {
	return &Error{Code: 2101, Message: "账号不存在"}
}

// StatusErrorByAccount 账号状态异常
func StatusErrorByAccount() *Error {
	return &Error{Code: 2102, Message: "账号状态异常"}
}

// ExistByAccount 名称已被占用
func ExistByAccount() *Error {
	return &Error{Code: 2103, Message: "账号已被占用"}
}

// ExistByName 名称已被占用
func ExistByName() *Error {
	return &Error{Code: 2104, Message: "名称已被占用"}
}

// IimitCountByRoleInCompany 指定角色数量超出限制
func IimitCountByRoleInCompany() *Error {
	return &Error{Code: 2105, Message: "指定角色数量超出限制"}
}

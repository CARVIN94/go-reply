package reply

// IimitCountInCompany 错误
func IimitCountInCompany() *Error {
	return &Error{Code: 3001, Message: "已达到可注册店铺数量上限"}
}

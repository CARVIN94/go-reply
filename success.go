package reply

// Success 执行成功
func Success() *Error {
	return &Error{Code: 0, Message: "执行成功"}
}

// CreateSuccess 新增成功
func CreateSuccess() *Error {
	return &Error{Code: 0, Message: "新增成功"}
}

// UpdateSuccess 更新成功
func UpdateSuccess() *Error {
	return &Error{Code: 0, Message: "更新成功"}
}

// DeleteSuccess 删除成功
func DeleteSuccess() *Error {
	return &Error{Code: 0, Message: "删除成功"}
}

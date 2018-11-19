package reply

// Object 数据返回格式
type Object struct {
	Code   int         `json:"code"`
	Object interface{} `json:"object"`
}

// Array 数据返回格式
type Array struct {
	Code  int         `json:"code"`
	Rows  interface{} `json:"rows"`
	Total int         `json:"total"`
}

// Error 错误返回格式
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

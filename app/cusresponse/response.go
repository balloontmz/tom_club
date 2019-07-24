package cusresponse

// ResponseFmt 返回格式结构体
type ResponseFmt struct {
	Ret  int         `json:"ret" xml:"ret"`
	Msg  string      `json:"msg" xml:"msg"`
	Data interface{} `json:"data" xml:"data"`
}

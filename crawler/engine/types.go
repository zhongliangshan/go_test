package engine

// 获得的请求集
type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}

// 结果集
type ParserResult struct {
	Items []interface{}
	Requests []Request
}

func NilRequest([]byte) ParserResult {
	return ParserResult{}
}

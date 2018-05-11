package engine

// 解析的参数 , 一个是url  一个是解析的函数(返回时 另外一个解析的参数集合和结果)
type Request struct {
	Url string
	ParserFunc func([]byte) ParserResult
}

type ParserResult struct {
	Requests []Request
	Items []Item // 定义接口表示随便返回什么值都可以
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
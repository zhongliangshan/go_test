package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParserRequest,
}

type ParserRequest struct {
	Item []interface{}
	Requests []Request
}

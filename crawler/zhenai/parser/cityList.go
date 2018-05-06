package parser

import (
	"github.com/zhongliangshan/test/crawler/engine"
	"regexp"
)

func ParserCityList(all []byte) engine.ParserResult {
	re := regexp.MustCompile(`<a\s+href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(all, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
	}
	return result

}

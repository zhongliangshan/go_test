// 城市解析器
package parser

import (
	"github.com/zhongliangshan/test/crawler2/engine"
	"regexp"
)

var (
	cityRe     = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityMoreRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParserCity(all []byte) engine.ParserResult {
	matches := cityRe.FindAllSubmatch(all, -1)

	result := engine.ParserResult{}

	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParserResult {
				return ParserProfile(bytes, string(m[1]), string(m[2]))
			},
		})

	}

	matches = cityMoreRe.FindAllSubmatch(all, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserCity,
		})
	}

	return result
}

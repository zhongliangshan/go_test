// 城市解析器
package parser

import (
	"github.com/zhongliangshan/test/crawler2/engine"
	"regexp"
)

var (
	cityRe     = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityMoreRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"[^>]*>([^<]+)</a>`)
)

func ParserCity(all []byte) engine.ParserResult {
	matches := cityRe.FindAllSubmatch(all, -1)

	result := engine.ParserResult{}

	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(bytes []byte) engine.ParserResult {
				return ParserProfile(bytes, name)
			},
		})

	}

	cityMoreMatches := cityMoreRe.FindAllSubmatch(all, -1)

	for _, cMM := range cityMoreMatches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(cMM[1]),
				ParserFunc: ParserCity,
			})

	}
	return result
}

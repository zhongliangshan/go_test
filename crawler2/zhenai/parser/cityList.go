package parser

import (
	"regexp"
	"github.com/zhongliangshan/test/crawler2/engine"
)

const cityListRe  =  `<a\s+href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParserCityList(all []byte) engine.ParserResult{
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(all , -1)

	result := engine.ParserResult{}
	limit := 10
	for _ , m := range matches {

		result.Items = append(result.Items , "City:" + string(m[2]))

		result.Requests = append(result.Requests , engine.Request{
			Url:string(m[1]),
			ParserFunc:ParserCity,
		})

		limit--
		if limit <=0 {
			break
		}

	}
	return result
}


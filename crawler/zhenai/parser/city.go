// 城市解析器
package parser

import (
<<<<<<< HEAD
	"regexp"
	"github.com/zhongliangshan/test/crawler/engine"
)

const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(all []byte) engine.ParserResult{
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(all , -1)

	result := engine.ParserResult{}

	for _ , m := range matches {

		result.Items = append(result.Items , "User:" + string(m[2]))

		result.Requests = append(result.Requests , engine.Request{
			Url:string(m[1]),
			ParserFunc:ParserProfile,
=======
	"github.com/zhongliangshan/test/crawler/engine"
	"regexp"
)

var (
	cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
)

func ParserCity(all []byte) engine.ParserResult {
	matches := cityRe.FindAllSubmatch(all, -1)

	result := engine.ParserResult{}

	for _, m := range matches {

		result.Items = append(result.Items, "User:"+string(m[2]))

		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParserProfile,
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
		})

	}
	return result
}
<<<<<<< HEAD

=======
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245

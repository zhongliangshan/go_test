package parser

import (
	"regexp"
	"fmt"
	"github.com/zhongliangshan/test/crawler/engine"
)

func ParserCityList(all []byte) engine.Request{
	re := regexp.MustCompile(`<a\s+href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(all , -1)
	request := engine.Request{}
	for _ , m := range matches {
		fmt.Printf("URL:%s  Citye: %s\n" , m[1] ,m[2])
	}

	fmt.Println(len(matches))

}

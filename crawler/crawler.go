package main

import (
	//"golang.org/x/text/encoding/simplifiedchinese"
	"github.com/zhongliangshan/test/crawler/engine"
	"github.com/zhongliangshan/test/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParserCityList,
	})

}
//	resp , err :=  http.Get("http://www.zhenai.com/zhenghun")
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//	if resp.StatusCode != http.StatusOK {
//		fmt.Println("Error:err http code : " , resp.StatusCode)
//		return
//	}
//
//	e := determineEncoding(resp.Body)
//	// 默认就是转化成utf8
//	//utf8Reader := transform.NewReader(resp.Body , simplifiedchinese.GBK.NewDecoder())
//	utf8Reader := transform.NewReader(resp.Body , e.NewDecoder())
//	all, err := ioutil.ReadAll(utf8Reader)
//	if err != nil {
//		panic(err)
//	}
//	//fmt.Printf("%s\n" , all)
//	getCityAndUrl(all)
//}
//
//func getCityAndUrl(all []byte) {
//	re := regexp.MustCompile(`<a\s+href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
//	matches := re.FindAllSubmatch(all , -1)
//
//	for _ , m := range matches {
//		fmt.Printf("URL:%s  Citye: %s\n" , m[1] ,m[2])
//	}
//
//	fmt.Println(len(matches))
//
//}
//
//// 探测html是什么编码格式
//func determineEncoding(r io.Reader) encoding.Encoding {
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		panic(err)
//	}
//	e, _, _ := charset.DetermineEncoding(bytes, "")
//
//	return e
//}

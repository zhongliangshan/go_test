package engine

import (
	"fmt"
	"log"

	"test/crawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request

	for _ , r := range seeds{
		requests = append(requests , r)
	}

	for len(requests) > 0 {
		r :=  requests[0]
		requests = requests[1:]
		log.Printf("Url:%s" , r.Url)
		body, err := fetcher.Fetch(r.url)
		if err != nil {
			log.Printf("Fetcheing error:url:%s error:%v" , r.Url , err)
			continue
		}
		parserResult := r.ParserFunc(body)
		requests = append(requests , parserResult.Requests...)

		for _,item := range parserResult.Item {
			log.Printf("Got item %s" , item)
		}
	}
}

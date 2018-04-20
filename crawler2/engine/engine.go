package engine

import (
	"github.com/zhongliangshan/test/crawler2/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request

	for _ , r:=range seeds {
		requests = append(requests, r)
	}



	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("start parser Url:%s" , r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher Url: %s error : %v" , r.Url , err)
			continue
		}

		parserResult := r.ParserFunc(body)
		requests = append(requests , parserResult.Requests...)

		for _ , item := range parserResult.Items {
			log.Printf("Got Item :%s" , item)
		}
	}
}

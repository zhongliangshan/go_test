package engine

import (
	"github.com/zhongliangshan/test/crawler/fetcher"
	"log"
)

type SimpleScheduler struct {
}

func (e SimpleScheduler) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}

func Worker(r Request) (ParserResult, error) {
	log.Printf("Url:%s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcheing error:url:%s error:%v", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), nil
}

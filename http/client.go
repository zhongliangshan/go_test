package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {

	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent" ,
		"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Mobile Safari/537.36")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect:%v"  ,req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	res, err := httputil.DumpResponse(resp , true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s" , res)
}

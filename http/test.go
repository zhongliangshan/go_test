package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"net/http/httputil"
	"fmt"
	"strconv"
)

type PostData struct {
	Url     string
	Param   map[string]string
	Timeout int
}

func (p *PostData) PostAlarm() (string, error) {

	data, err := json.Marshal(p.Param)

	if err != nil {
		return "", err
	}

	request, err := http.NewRequest(http.MethodPost, p.Url, strings.NewReader(string(data)))
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type",
		"application/json; charset=utf-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	res, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func main() {
	fmt.Println(strconv.FormatFloat(float64(0)/float64(1), 'f', 2, 64))
}

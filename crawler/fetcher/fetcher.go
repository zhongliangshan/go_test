package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
<<<<<<< HEAD
	"golang.org/x/text/encoding/unicode"
	"time"
=======
	"net/http"
>>>>>>> fd473d8fc6f664aece8abcb043166579d4fd1245
)
var Rartelimiter = time.Tick(time.Millisecond * 100)
func Fetch(url string) ([]byte, error) {
	<-Rartelimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error:err http code : %d", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	// 默认就是转化成utf8
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// 探测html是什么编码格式
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error :  %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}

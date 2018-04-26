package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"io"
	"bufio"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
	"log"
	"time"
)


var limitRater = time.Tick(time.Millisecond * 100)
func Fetch(url string)  ([]byte, error) {
	resp , err :=  http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil , fmt.Errorf("Error:err http code : %s" , resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	// 默认就是转化成utf8
	//utf8Reader := transform.NewReader(resp.Body , simplifiedchinese.GBK.NewDecoder())
	utf8Reader := transform.NewReader(resp.Body , e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}


// 探测html是什么编码格式
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("determineEncoding error: %v" , err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}

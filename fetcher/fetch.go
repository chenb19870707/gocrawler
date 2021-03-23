package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)
var ratelimit = time.Tick(2000*time.Microsecond)
func Fetch(url string) ([] byte,error)  {
	<-ratelimit

	client := &http.Client{}
	req,err := http.NewRequest("GET",url,nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv2.0.1) Gecko/20100101 Firefox/4.0.1")

	res,err := client.Do(req)
	if err !=  nil{
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error status code %d",res.StatusCode)
	}

	bodyReader := bufio.NewReader(res.Body)
	e := DetectEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())

	return  ioutil.ReadAll(utf8Reader)
}

func DetectEncoding(r *bufio.Reader)  encoding.Encoding{
	bytes,err := r.Peek(1024)
	if err != nil{
		log.Printf("fetch err:%v",err)
		return unicode.UTF8
	}

	e,_,_ := charset.DetermineEncoding(bytes,"")

	return e
}
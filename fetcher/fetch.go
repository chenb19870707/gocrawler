package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"net/url"
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
		fmt.Printf("Error status code %d\n",res.StatusCode)
		return nil,errors.New("Error status code ")
	}

	bodyReader := bufio.NewReader(res.Body)
	e := DetectEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())

	return  ioutil.ReadAll(utf8Reader)
}

//http代理访问访问
func ProxyFetch(weburl string )([]byte,error){

	<-ratelimit


	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:7890")//根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	req,err:= http.NewRequest("GET",weburl,nil)
	if err!=nil{
		return nil,fmt.Errorf("ERROR: get url:%s",weburl)
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	resp,err:= client.Do(req)
	if err!=nil{
		return nil,fmt.Errorf("ERROR: get url:%s",weburl)
	}
	bodyReader:= bufio.NewReader(resp.Body)
	e:= DetectEncoding(bodyReader)
	if err!=nil{
		return nil,fmt.Errorf("ERROR: get url:%s",weburl)
	}
	utf8Reader:= transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func DetectEncoding(r *bufio.Reader)  encoding.Encoding{
	bytes,err := r.Peek(1024)
	if err != nil{
		//log.Printf("fetch err:%v",err)
		return unicode.UTF8
	}

	e,_,_ := charset.DetermineEncoding(bytes,"")

	return e
}
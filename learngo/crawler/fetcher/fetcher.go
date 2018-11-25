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
	"net/http"
)


//获取url中的数据（utf-8）
func Fetcher(url string)([]byte,error){
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println()
		return nil,fmt.Errorf("Error: status code %d", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	//借助gopm get -g -v golang.org/x/text中的transform.NewReader将编码转为utf-8
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())

	//可以直接返回
	return ioutil.ReadAll(utf8Reader)

}

func Fetcher2(url string)([]byte,error){
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("Fetch error :%v",err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Fetch error :%v",err)
	}

	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println()
		return nil,fmt.Errorf("Error: status code %d", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)
	//借助gopm get -g -v golang.org/x/text中的transform.NewReader将编码转为utf-8
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())

	//可以直接返回
	return ioutil.ReadAll(utf8Reader)

}

/*
 借助gopm get -g -v golang.org/x/net/html
 中的charset.DetermineEncoding,来判断当前的内容是什么编码
 */
func determineEncoding(r io.Reader) encoding.Encoding{
	bytes,err:=bufio.NewReader(r).Peek(1024)
	if err != nil {
		fmt.Printf("Fetcher error %v",err)
		//最好不要panic
		return unicode.UTF8
	}

	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

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

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code %d", resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)
	//借助gopm get -g -v golang.org/x/text中的transform.NewReader将编码转为utf-8
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", all)

}

/*
 借助gopm get -g -v golang.org/x/net/html
 中的charset.DetermineEncoding,来判断当前的内容是什么编码
 */
func determineEncoding(r io.Reader) encoding.Encoding{
	bytes,err:=bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

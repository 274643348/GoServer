package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func main() {

	//fmt.Println("asdfasdfg---%d",findMaxLenOfNoRepeat("asdfasdfg"))

	fmt.Println("aabcdab过程中---%d",findMaxLenOfNoRepeat("aabcdab过程中"))

}

func findMaxLenOfNoRepeat(message string) int{
	charMap:= make(map[rune]int)
	maxLen:=0
	start:=0
	for i,ch:=range []rune(message)  {
		if index,err :=charMap[ch]; err && index >=start{
			start =index+1
		}

		if i - start +1 > maxLen{
			maxLen = i - start +1
		}
		charMap[ch] = i
	}
	return maxLen

}

func chineseRead(){
	s:="Yes司马懿"
	for _,b:=range []byte(s){
		fmt.Printf("%X ",b)
	}
	fmt.Println("\n")
	for i,ch:=range s{
		fmt.Printf("(%d,%X) ",i,ch)
	}
	fmt.Println("\n")
	fmt.Printf("utf8-count--%d\n",utf8.RuneCountInString(s))
	fmt.Printf("s®-len--%d\n",len(s))

	fmt.Println("\n")
	bytes:=[] byte(s)
	for len(bytes)>0  {
		ch,size := utf8.DecodeRune(bytes)
		fmt.Printf("%c--%d\n",ch,size)
		bytes =bytes[size:]
	}

	for i,ch:=range []rune(s){
		fmt.Printf("%d---%c\n",i,ch)
	}

	fmt.Printf("%d",strings.EqualFold("aa","aa"))
}



func sizTypesizeofeof() {
	var i int=1
	var i2 int8 = 2
	var i3 int16  = 3
	var i4 int32  = 4
	var i5 int64  = 5


	println(unsafe.Sizeof(i))
	println(unsafe.Sizeof(i2))
	println(unsafe.Sizeof(i3))
	println(unsafe.Sizeof(i4))
	println(unsafe.Sizeof(i5))

}

func typeIota(){
	const(
		i=3<<iota
		b=2<<iota
		c
		d
	)
	println(i,b,c,d)
}

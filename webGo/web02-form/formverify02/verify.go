package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	http.HandleFunc("/login",login)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		panic(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		 t,err :=template.ParseFiles("./web02-form/formverify02/login.gtpl")
		if err != nil {
			panic(err)
			return
		}
		 err = t.Execute(w,nil)
		if err != nil {
			panic(err)
			return
		}
	}else{
		////必填检查
		//if err := verifyMustFillIn(r,"username"); err!=nil {
		//	fmt.Fprintln(w,err.Error())
		//	return
		//}
		//if err := verifyMustFillIn(r,"password"); err!=nil {
		//	fmt.Fprintln(w,err.Error())
		//	return
		//}
		//
		//
		////数字检查
		//if err := verifyMustNum(r,"number"); err!=nil {
		//	fmt.Fprintln(w,err.Error())
		//	return
		//}
		////数字检查(正则表达式)
		//if err := verifyRegexpMustNum(r,"number"); err!=nil {
		//	fmt.Fprintln(w,err.Error())
		//	return
		//}


		////中文检查(正则表达式)
		//if err := verifyRegexpMustChinese(r,"zhongwen"); err!=nil {
		//	fmt.Fprintln(w,err.Error())
		//	return
		//}

		////英文检查(正则表达式)
		//if err := verifyRegexpMustEnglish(r,"yingwen"); err!=nil {
		//	fmt.Fprintln(w,err.Error())
		//	return
		//}

		//英文检查(正则表达式)
		if err := verifyRegexpMustPhone(r,"phoneNum"); err!=nil {
			fmt.Fprintln(w,err.Error())
			return
		}



		fmt.Println(r.Form.Get("username"))
		fmt.Println(r.Form.Get("password"))

		fmt.Fprintln(w,r.Form.Get("username"))
		fmt.Fprintln(w,r.Form.Get("password"))
	}
}

//必填字段检查
func verifyMustFillIn(r *http.Request,key string)error{
	value := r.Form.Get(key)
	if len(value) == 0  {
		fmt.Println("error："+key+" is empty")
		return errors.New("error："+key+" is empty")
	}
	return  nil
}


//数字字段检查
func verifyMustNum(r *http.Request,key string)error{
	err :=verifyMustFillIn(r,key)
	if err != nil{
		return  err
	}

	getint,err :=  strconv.Atoi(r.Form.Get(key))
	if err != nil  {
		fmt.Println(err)
		return err
	}

	fmt.Println(key +" is " + fmt.Sprint(getint))
	return  nil
}

//数字字段检查(正则表达式)
func verifyRegexpMustNum(r *http.Request,key string)error{
	err :=verifyMustFillIn(r,key)
	if err != nil{
		return  err
	}

	matched, err :=  regexp.MatchString(`^[0-9]+$`,r.Form.Get(key))
	if err != nil ||  !matched{
		fmt.Println(err)
		return err
	}
	fmt.Println(key +" is number ?" + fmt.Sprint(matched))
	return  nil
}

//中文字段检查(正则表达式)
func verifyRegexpMustChinese(r *http.Request,key string)error{
	err :=verifyMustFillIn(r,key)
	if err != nil{
		return  err
	}

	matched, err :=  regexp.MatchString(`^\p{Han}+$`,r.Form.Get(key))

	if err != nil || !matched {
		fmt.Println("error："+key+" is err")
		return errors.New("error："+key+" is err")
	}
	fmt.Println(key +" is chinese " + fmt.Sprint(r.Form.Get(key)))
	return  nil
}

//英文字段检查(正则表达式)
func verifyRegexpMustEnglish(r *http.Request,key string)error{
	err :=verifyMustFillIn(r,key)
	if err != nil{
		return  err
	}

	matched, err :=  regexp.MatchString(`^[a-zA-Z]+$`,r.Form.Get(key))
	if err != nil || !matched {
		fmt.Println("error："+key+" is err")
		return errors.New("error："+key+" is err")
	}
	fmt.Println(key +" is english " + fmt.Sprint(r.Form.Get(key)))
	return  nil
}

//手机地址字段检查(正则表达式)---有时间实现
func verifyRegexpMustPhone(r *http.Request,key string)error{
	err :=verifyMustFillIn(r,key)
	if err != nil{
		return  err
	}

	matched, err :=  regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4-8})$`,r.Form.Get(key))
	if err != nil || !matched {
		fmt.Println("error："+key+" is err")
		return errors.New("error："+key+" is err")
	}
	fmt.Println(key +" is phoneNum " + fmt.Sprint(r.Form.Get(key)))
	return  nil
}
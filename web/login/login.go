package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("以下打印sayhelloName请求的信息")
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("以下打印login请求的信息")
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		fmt.Println("login success")
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		v := url.Values{}
		v.Set("name", "Ava")
		v.Add("friend", "Jess")
		v.Add("friend", "Sarah")
		v.Add("friend", "Zoe")
		// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		fmt.Println(v.Get("name"))
		fmt.Println(v.Get("friend"))
		fmt.Println(v["friend"])
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("以下打印index请求的信息")
	if r.Method == "GET" {
		template.ParseFiles("index.html")
	}
}

func main() {
	http.HandleFunc("/", index)                    //设置访问的路由
	http.HandleFunc("/sayhelloName", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)               //设置访问的路由
	err := http.ListenAndServe(":9090", nil)       //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package main

import(
    "fmt"
    "net/http"
    "strings"
    "log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request){ //http.w写入，http.r读取
    r.ParseForm() //解析参数，默认是不会解析的
    fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form { 
        //当url为localhost:8080/?url_long=111&url_long=222   输出为key=url_long val=111222
        fmt.Println("key:",k)
        fmt.Println("val:",strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") //这个是写入到w的是输出到客户端的
}

func main(){
    http.HandleFunc("/", sayhelloName) //设置访问的路由
    err := http.ListenAndServe(":8080",nil) //设置监听端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

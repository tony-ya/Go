package main

import (
    "fmt"
    "crypto/md5"
    "crypto/sha1"
    "io"
    "flag"
)

var (
    str = flag.String("s", "", "str of the string")
)

func a(data string) string{
    t := md5.New()
    io.WriteString(t,data)
    return fmt.Sprintf("%x",t.Sum(nil))
}

func b(data string) string {
    t := sha1.New()
    io.WriteString(t, data)
    return fmt.Sprintf("%x", t.Sum(nil))
}

func main(){
    flag.Parse()
    var str2 = *str
    fmt.Printf("MD5 : %s\n", a(str2))
    fmt.Printf("SHA1 : %s\n", b(str2))
}

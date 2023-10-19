package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("HANDLING URL IN GOLANG")
	myurl := "https://lco.dev:3000/learn?coursename=Golang&paymentid=12345"
	result, _ := url.Parse(myurl)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Scheme)
	fmt.Println(result.Query())
	fmt.Println(result.Query()["paymentid"])
	fmt.Println(result.OmitHost)
}

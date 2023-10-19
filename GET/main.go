package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev/"

func main() {
	fmt.Println(" GET IN GOLANG")
	response, _ := http.Get(url)
	data, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Print(string(data))
}

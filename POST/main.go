package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "https://lco.dev/"

func main() {
	content := strings.NewReader("HELLO THIS IS CONTENT")
	fmt.Println(" POST IN GOLANG")
	response, _ := http.Post(url, "applications/json", content)
	data, _ := io.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Print(string(data))

}

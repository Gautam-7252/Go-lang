package main

func main() {

	// Switch case

	// var a int = 5
	// switch a {
	// case 1:
	// 	fmt.Println("The value is 1: ")
	// case 2:
	// 	fmt.Println("The value is 2: ")
	// case 5:
	// 	fmt.Println("The value is 5: ")
	// case 7:
	// 	fmt.Println("The value is 7: ")
	// case 11:
	// 	fmt.Println("The value is 11: ")
	// }

	//Float to String

	// f := 3.14159
	// s := strconv.FormatFloat(f, 'g', 5, 32)
	// fmt.Printf("Type of s: %T\n", s)
	// fmt.Printf("Type of f: %T\n", f)
	// fmt.Println()
	// fmt.Printf("Value of s: %v\n", s)
	// fmt.Printf("Value of f: %v", f)

	//Array

	// a := [5]int{10, 20, 30, 40, 50}
	// b := a[0:4]
	// fmt.Printf("%T %T", a, b)

	//Working with File

	// func main() {
	// 	fmt.Println("Working with Files in Golang")
	// 	file, err := os.Create("./test.txt")
	// 	content := "HELLO"
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	length, err := io.WriteString(file, content)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("The length is :", length)
	// 	readFile("./test.txt")
	// 	defer file.Close()
	// }
	// func readFile(filename string) {
	// 	databyte, err := os.ReadFile(filename)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Databyte : ", string(databyte))
	// }

	//Handling Web Request

	// 	const URL = "http://lco.dev"
	// func main() {
	// 	fmt.Println("Handling Web Request")
	// 	response, _ := http.Get(URL)
	// 	fmt.Println(response)
	// 	Data, _ := ioutil.ReadAll(response.Body)
	// 	fmt.Println("DATA : ", string(Data))
	// 	defer response.Body.Close()
	//}

	//Handling URL

	// const myurl string = "https://lco.in:3000/learn?coursename=reactJS&paymentid=12345"
	// func main() {
	// 	result, _ := url.Parse(myurl)
	// 	fmt.Println(result.Scheme)
	// 	fmt.Println(result.Host)
	// 	fmt.Println(result.Path)
	// 	fmt.Println(result.Port())
	// 	fmt.Println(result.Query()["coursename"])
	// 	partsofmyurl := &url.URL{
	// 		Scheme: "https",
	// 		Host:   "lco.dev",
	// 		Path:   "/tutcss",
	// 	}
	// 	anotherurl := partsofmyurl.String()
	// 	fmt.Println(anotherurl)
	// }

}

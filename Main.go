package main

import Console "fmt"
import (
	Runtime "runtime"
	"./Parser"
	Strings "strings"
	Strconv "strconv"
)

func main() {
	// Multiple CPU Processing
	Runtime.GOMAXPROCS(4)

	Console.Println("[Nicname Parser]")
	Console.Print("사이트 주소(쉼표로 구분) : ")
	var inputData string
	Console.Scan(&inputData)
	var urls []string = Strings.Split(inputData, ",")

	var i int
	for i = 0; i < 100; i++ {
		urls = append(urls, "https://sir.kr/qa/p" + Strconv.Itoa(i) + "?s_tag=%EA%B7%B8%EB%88%84%EB%B3%B4%EB%93%9C5")
	}

	for _, url := range urls {
		Console.Println(url)
		Parser.GetNickName(Parser.GetHtmlBody(url))
	}

}
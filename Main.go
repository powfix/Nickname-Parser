package main

import Console "fmt"
import (
	Runtime "runtime"
	"./Parser"
	//Strings "strings"
	Strconv "strconv"
	Sync "sync"
)

var resultData = Parser.ResultData{}
func main() {
	// Multiple CPU Processing
	Runtime.GOMAXPROCS(4)

	//Console.Println("[Nicname Parser]")
	//Console.Print("사이트 주소(쉼표로 구분) : ")
	//var inputData string
	//Console.Scan(&inputData)

	var cnt int = 0
	var taskLimit int = 10
	var wait Sync.WaitGroup
	wait.Add(taskLimit)
	for i := 1; i < 60; i++ {
		cnt++
		go func() {
			defer wait.Done()
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/qa/p" + Strconv.Itoa(i) + "?unanswered"))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}

	Console.Println("※ 작업이 완료되었습니다")
	resultData.RemoveDuplicate()
	Parser.PrintResult(&resultData)
}
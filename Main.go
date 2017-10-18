package main

import Console "fmt"
import (
	Runtime "runtime"
	"./Parser"
	//Strings "strings"
	Strconv "strconv"
	Sync "sync"
	FIle "./File"
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
	var taskLimit int = 4
	var wait Sync.WaitGroup
	wait.Add(taskLimit)
	for i := 1; i < 120; i++ {
		cnt++
		go func() {
			defer wait.Done()
			//tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/qa/p" + Strconv.Itoa(i) + "?unanswered"))
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/talk/p" + Strconv.Itoa(i)))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}
	for i := 1; i < 4758; i++ {
		cnt++
		go func() {
			defer wait.Done()
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/cm_free/p" + Strconv.Itoa(i)))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}
	for i := 1; i < 2626; i++ {
		cnt++
		go func() {
			defer wait.Done()
			//tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/qa/p" + Strconv.Itoa(i) + "?unanswered"))
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/qa/p" + Strconv.Itoa(i)))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}
	for i := 1; i < 403; i++ {
		cnt++
		go func() {
			defer wait.Done()
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/qa/p" + Strconv.Itoa(i) + "?s_tag=%EC%98%81%EC%B9%B4%ED%8A%B85"))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}
	for i := 1; i < 8; i++ {
		cnt++
		go func() {
			defer wait.Done()
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/yc5_tip/p" + Strconv.Itoa(i)))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}
	for i := 1; i < 91; i++ {
		cnt++
		go func() {
			defer wait.Done()
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/cm_bug/p" + Strconv.Itoa(i) + "?sca=%EC%98%81%EC%B9%B4%ED%8A%B85"))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}
	for i := 1; i < 457; i++ {
		cnt++
		go func() {
			defer wait.Done()
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/cm_humor/p" + Strconv.Itoa(i)))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}
	for i := 1; i < 20; i++ {
		cnt++
		go func() {
			defer wait.Done()
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/cm_quiz/p" + Strconv.Itoa(i)))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}
	for i := 1; i < 27; i++ {
		cnt++
		go func() {
			defer wait.Done()
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/cm_adsense/p" + Strconv.Itoa(i)))
			resultData.Append(&tmp)
		}()
		if cnt > taskLimit {
			wait.Wait()
			cnt = 0
			wait.Add(taskLimit)
		}
		Console.Printf("# Go routine start : %04d\n", i)
	}

	for i := 1; i < 1021; i++ {
		cnt++
		go func() {
			defer wait.Done()
			tmp := Parser.GetNickName(Parser.GetHtmlBody("https://sir.kr/cm_free_10y/p" + Strconv.Itoa(i)))
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
	FIle.WriteIntoFile(FIle.SliceStringToCsvString(resultData.GetResult()))
}
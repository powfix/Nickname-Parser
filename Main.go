package main

import Console "fmt"
import Runtime "runtime"

func main() {
	// Multiple CPU Processing
	Runtime.GOMAXPROCS(4)

	urls := []string{}
	urls = append(urls, "https://sir.kr/qa/?s_tag=%EA%B7%B8%EB%88%84%EB%B3%B4%EB%93%9C5")
	urls = append(urls, "https://sir.kr/qa/p2?s_tag=%EA%B7%B8%EB%88%84%EB%B3%B4%EB%93%9C5")
	urls = append(urls, "https://sir.kr/qa/p3?s_tag=%EA%B7%B8%EB%88%84%EB%B3%B4%EB%93%9C5")
	urls = append(urls, "https://sir.kr/qa/p4?s_tag=%EA%B7%B8%EB%88%84%EB%B3%B4%EB%93%9C5")
	urls = append(urls, "https://sir.kr/qa/p5?s_tag=%EA%B7%B8%EB%88%84%EB%B3%B4%EB%93%9C5")

	for _, url := range urls {
		// TODO: go 스레드 처리시 Console 출력안되는 문제해결
		//getNickName(url)

		go func() {
			Console.Println(url)
		}()
	}
}
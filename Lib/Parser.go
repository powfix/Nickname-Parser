package Lib

import Console "fmt"
import HTTP "net/http"
import Scrape "github.com/yhat/scrape"
import HTML "golang.org/x/net/html"
import Atom "golang.org/x/net/html/atom"
import Regexp "regexp"

type ResultData struct {
	result []string
	totalCount, includeCount, excludeCount int
}

func printResult(resultData *ResultData) {
	Console.Println("전체/추가됨/제외됨 : ", resultData.totalCount, "/", resultData.includeCount, "/", resultData.excludeCount)
}

func getHtmlBody(url string) *HTML.Node {
	resp, err := HTTP.Get(url)
	if err != nil {panic(err)}
	root, err := HTML.Parse(resp.Body)
	if err != nil {panic(err)}
	return root
}

func getNickName(root *HTML.Node) (resultData *ResultData) {
	// define a matcher
	matcher := func(n *HTML.Node) bool {
		// must check for nil values
		if n.DataAtom == Atom.Span && n.Parent != nil && n.Parent.Parent != nil {
			return Scrape.Attr(n, "class") == "member"
		}
		return false
	}
	articles := Scrape.FindAll(root, matcher)
	/** Count reset */
	resultData.totalCount, resultData.includeCount, resultData.excludeCount = 0, 0, 0
	for i, article := range articles {
		text := Scrape.Text(article)
		hasKorean, _ := Regexp.MatchString("[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]", text)		// 한글이 포함되어 있는가
		hasEnglish, _ := Regexp.MatchString("[a-z A-Z]", text)					// 영문이 포함되어 있는가
		hasNumber, _ := Regexp.MatchString("[0-9]", text)						// 숫자가 포함되어 있는가
		if hasKorean && hasEnglish && !hasNumber {										// 한글이 포함되어 있고 영문과 숫자가 포함되어 있지 않을 경우
			Console.Printf("%2d %s (%s)\n", i, text, Scrape.Attr(article, "class")) // Console 표시 ex) "0 민쯩먼저깔게요 (member)"
			resultData.result = append(resultData.result, Scrape.Text(article))
			resultData.includeCount++
		} else {
			resultData.excludeCount++
		}
		resultData.totalCount++
	}
	printResult(resultData)
	return
}

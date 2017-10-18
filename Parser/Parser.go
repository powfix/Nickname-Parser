package Parser

import Console "fmt"
import HTTP "net/http"
import Scrape "github.com/yhat/scrape"
import HTML "golang.org/x/net/html"
import Atom "golang.org/x/net/html/atom"
import (
	Regexp "regexp"
	"unicode"
)

type ResultData struct {
	result []string
	totalCount, includeCount, excludeCount int
}
func itemCounter() func() int {
	itemCount := 0
	return func() int {
		itemCount += 1
		return itemCount
	}
}


func PrintResult(resultData *ResultData) {
	Console.Println("전체/추가됨/제외됨 : ", resultData.totalCount, "/", resultData.includeCount, "/", resultData.excludeCount, '\n')
}

func GetHtmlBody(url string) *HTML.Node {
	resp, err := HTTP.Get(url)
	if err != nil {panic(err)}
	root, err := HTML.Parse(resp.Body)
	if err != nil {panic(err)}
	return root
}

func GetNickName(root *HTML.Node) ResultData {
	// define a matcher
	matcher := func(node *HTML.Node) bool {
		// must check for nil values
		if node.DataAtom == Atom.Span && node.Parent != nil && node.Parent.Parent != nil {
			return Scrape.Attr(node, "class") == "member"
		}
		return false
	}
	articles := Scrape.FindAll(root, matcher)

	/** Count reset */
	var resultData = ResultData{totalCount: 0, includeCount: 0, excludeCount: 0}
	for i, article := range articles {
		text := Scrape.Text(article)
		hasKorean, _ := Regexp.MatchString("[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]", text)		// 한글이 포함되어 있는가
		hasEnglish, _ := Regexp.MatchString("[a-z A-Z]", text)					// 영문이 포함되어 있는가
		hasNumber, _ := Regexp.MatchString("[0-9]", text)						// 숫자가 포함되어 있는가
		if hasKorean && !hasEnglish && !hasNumber {										// 한글이 포함되어 있고 영문과 숫자가 포함되어 있지 않을 경우
			Console.Printf("%03d %s (%s)\n", i, text, Scrape.Attr(article, "class")) // Console 표시 ex) "0 민쯩먼저깔게요 (member)"
			resultData.result = append(resultData.result, Scrape.Text(article))
			resultData.includeCount++
		} else {
			resultData.excludeCount++
		}
		resultData.totalCount++
	}
	PrintResult(&resultData)
	return resultData
}
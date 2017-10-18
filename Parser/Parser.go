package Parser

import Console "fmt"
import HTTP "net/http"
import Scrape "github.com/yhat/scrape"
import HTML "golang.org/x/net/html"
import Atom "golang.org/x/net/html/atom"
import Regexp "regexp"

type ResultData struct {							// 결괏값 구조체
	result []string							// 닉네임 배열
	totalCount, includeCount, excludeCount int		// 전체 닉네임 카운트, 포함된 닉네임 카운트, 제외된 닉네임 카운트
}
type Result interface {
	GetResult() []string		// 닉네임 배열
	GetTotalCount() int			// 전체 닉네임 카운트
	GetIncludeCount() int		// 정규식 조건에서 포함된 닉네임 카운트
	GetExcludeCount() int		// 정규식 조건에서 제외된 닉네임 카운트
	GetItemCount() int			// 실제 닉네임 요소 카운트
	append()					// 슬라이스 추가
}
func (resultData ResultData) GetResult() []string {
	return resultData.result
}
func (resultData ResultData) GetTotalCount() int {
	return resultData.totalCount
}
func (resultData ResultData) GetIncludeCount() int {
	return resultData.includeCount
}
func (resultData ResultData) GetExcludeCount() int {
	return resultData.excludeCount
}
func (resultData ResultData) GetItemCount() int {
	return len(resultData.result)
}
func (resultData ResultData) append(data ResultData) ResultData {
	resultData.result = append(resultData.result, data.result...)
	resultData.totalCount += data.totalCount
	resultData.includeCount += data.includeCount
	resultData.excludeCount += data.excludeCount
	return resultData
}
func (resultData ResultData) RemoveDuplicate() ResultData {
	resultData.result = removeDuplicates(resultData.result)
	return resultData
}

var TresultData = ResultData{}

func PrintResult(resultData *ResultData) {
	for i, text := range resultData.result {
		Console.Printf("%03d %s\n", i, text) // Console 표시 ex) "27 은하수"
	}
	Console.Printf("※ 결괏값 : 전체(%d), 추가(%d), 정규식 제외(%d), 중복 제외(%d)\n\n", resultData.GetTotalCount(), resultData.GetIncludeCount(),resultData.GetExcludeCount(), resultData.includeCount - len(resultData.result))
}

func GetHtmlBody(url string) *HTML.Node {
	resp, err := HTTP.Get(url)
	if err != nil {
		Console.Println("Exception : ", err)
	}
	root, err := HTML.Parse(resp.Body)
	if err != nil {
		Console.Println("Exception : ", err)
	}
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
	for _, article := range articles {
		text := Scrape.Text(article)
		hasKorean, _ := Regexp.MatchString("[ㄱ-ㅎ|ㅏ-ㅣ|가-힣]", text)		// 한글이 포함되어 있는가
		hasEnglish, _ := Regexp.MatchString("[a-z A-Z]", text)					// 영문이 포함되어 있는가
		hasNumber, _ := Regexp.MatchString("[0-9]", text)						// 숫자가 포함되어 있는가
		if hasKorean && !hasEnglish && !hasNumber {										// 한글이 포함되어 있고 영문과 숫자가 포함되어 있지 않을 경우
			resultData.result = append(resultData.result, Scrape.Text(article))			// ResultData(Local) 변수의 슬라이스에 추가
			resultData.includeCount++
		} else {
			resultData.excludeCount++
		}
		resultData.totalCount++
	}
	TresultData.append(resultData)				// 전역변수에 추가
	PrintResult(&resultData)
	/** 코드 퍼포먼스를 위해 중복제거(removeDuplicates)는 마지막에 실행하도록 한다 */
	return resultData
}

func removeDuplicates(strings []string) []string {
	encountered := map[string]bool{}
	for i:= range strings {
		encountered[strings[i]] = true
	}
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}
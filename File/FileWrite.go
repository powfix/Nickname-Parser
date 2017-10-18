package File

import (
	"os"
	"time"
	"strconv"
)

func WriteIntoFile(str string)  {
	// 출력파일 생성
	now := time.Now()
	var fileName string = strconv.Itoa(now.Hour()) + "_" + strconv.Itoa(now.Minute()) + "_" + strconv.Itoa(now.Second()) + ".csv"
	fo, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	fo.WriteString(str)
}

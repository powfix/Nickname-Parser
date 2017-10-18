package Parser

// http://www.inven.co.kr/board/powerbbs.php?come_idx=1565&query=list&my=&category=&category2=&sort=PID&orderby=&name=&subject=&content=&keyword=&sterm=&eq=&iskin=&mskin=&p=284
// func someThing()  {
// 	for i = 1; i < 284; i++ {
// 		cnt++; totalCount++
// 		wait.Add(1)
// 		go func() {
// 			defer wait.Done()
// 			tmp := Parser.GetNickName(Parser.GetHtmlBody("http://www.inven.co.kr/board/powerbbs.php?come_idx=1565&query=list&my=&category=&category2=&sort=PID&orderby=&name=&subject=&content=&keyword=&sterm=&eq=&iskin=&mskin=&p=" + Strconv.Itoa(int(i))))
// 			resultData.Append(&tmp)
// 		}()
// 		if cnt > taskLimit {
// 			wait.Wait()
// 			cnt = 0
// 		}
// 		Console.Printf("# Go routine start : %04d\n", totalCount)
// 	}
// }
package main

import "fmt"

// 関数名はキャメルケース
// 関数名（引数 引数型）戻り型
func stringRetrun(arg string) string {
	return arg
}

func main() {
	fmt.Println(stringRetrun("Hello,World"))
}

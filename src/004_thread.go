package main

// 並列処理
import (
	"fmt"
	"time"
)

func say() {
	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
}

func main() {
	// 並列処理の開始
	go say()

	// メイン処理
	fmt.Println("5 second wait")
	time.Sleep(5000 * time.Millisecond)
	fmt.Println("wait end")

	// メインが終了すると並列処理も強制終了
}

package main

// 同期処理
import (
	"fmt"
	"time"
)

func say(c chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(i)
	}
	c <- true
}

func main() {
	channel := make(chan bool)
	go say(channel)
	fmt.Println("say end wait")
	b := <-channel
	fmt.Println(b)
	fmt.Println("say end")
}

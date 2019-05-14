package main

import (
	"fmt"
)

func main() {
	//channelFor()
	//channelFor()
	channelSelect()
}

func chanellSend() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Print(<-ch)
}

//channelに対してもrangeを利用してforを適用できる。
func channelFor() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	//closeされたらループを抜ける。
	for v := range ch {
		fmt.Print(v)
	}
}

//select文を利用することにより、複数のチャネルの待ち受け、送信処理を行うことができる。
func channelSelect() {

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch1 <- 1
	ch2 <- 2

	//slect文のcase節はチャネル関連の処理を伴う必要がある。
	//複数のcase節の処理が可能である場合（ブロックしないとき）、goランタイムは処理可能であるcase文をランダムに選択する。
	//以下の場合は、1つ目のcaseまたは2つ目のcaseのどちらかが、ランダムに選択される。
	select {
	case v := <-ch1:
		fmt.Printf("Recieved ch1: %d", v)
	case v := <-ch2:
		fmt.Printf("Recieved ch2: %d", v)
	//処理が可能であるcase節が存在しない時は、default節が実行される。
	default:
		fmt.Printf("Default")
	}
}

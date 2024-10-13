package main

import (
	"fmt"
	"time"
)
// 関数呼び出し
func main(){
	greeting()
}

// 挨拶
func greeting() {
	// 現在の日時を取得
	currentTime := time.Now()
	// 時間
	hour := currentTime.Hour()
	// Printlnで参照可能
	fmt.Println(currentTime)
	fmt.Println(hour)
	//時間ごとに挨拶を分ける
	if hour < 12 {
		fmt.Println("Good Morning")
	} else if hour < 18 {
		fmt.Println("Good Afternoon")
	} else {
		fmt.Println("Good Evening")
	}
}
package main

import ("fmt")

func main(){
	//  varを用いた例
	var num int
	num = 1
	fmt.Println(num)
	//  varを省略した例
	str := "文字"
	fmt.Println(str)
    //  デフォルトはfalseらしい
	var booleanVal bool = true
	fmt.Println(booleanVal)
}
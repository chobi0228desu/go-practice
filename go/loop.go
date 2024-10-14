package main

import ("fmt")

func main(){
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		// 0
		// 1
		// 2
		// 3
		// 4
	}

	for i, v := range []string{"foo", "bar", "baz"} {
		fmt.Println(i, v)
		// 0 foo
		// 1 bar
		// 2 baz
	}
}
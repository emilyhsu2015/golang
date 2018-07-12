package main

import (
	"fmt"
	"time"
)

func main() {
	var totle int = 0
	ch := make(chan int)
	// goroutine1
	for i := 1; i <= 20; i++ {
		go work(i, ch)
		totle = totle + <-ch
	}

	fmt.Println("ok", totle)
	time.Sleep(time.Second * 1)
}

func work(num int, c chan int) {
	var a int = 1
	for j := 1; j <= num; j++ {
		//fmt.Printf("階層%+v : a%+v , j%+v \n", num, a, j)
		if j != 1 {
			a = a * j
		} else {
			a = 1
		}
	}
	fmt.Printf("階層%+v 總和%+v \n", num, a)
	c <- a
}

package main

import (
	"fmt"
	"time"
)

//伐
type lumberjack struct {
	name  string  //姓名
	pulp  float32 //每次產出紙漿
	time  int     //每次製作需要時間
	color int     //訊息顯示的顏色
}

//造紙廠
type papermaking struct {
	name  string  //姓名
	paper int     //每次產出紙
	time  int     //每次製作需要時間
	pulp  float32 //每次需要花費的紙漿
	color int     //訊息顯示的顏色
}

//印刷廠
type printingfactory struct {
	name  string //姓名
	paper int    //每次產出紙
	time  int    //每次製作需要時間
	color int    //訊息顯示的顏色
}

func main() {
	//比賽結束訊息
	//var done string

	//伐
	l1 := lumberjack{
		name:  "伐木廠",
		pulp:  1,
		time:  1,
		color: 31,
	}

	//造紙廠1
	p1 := papermaking{
		name:  "造紙廠1",
		paper: 5000,
		time:  1,
		pulp:  0.5,
		color: 32,
	}
	p2 := papermaking{
		name:  "造紙廠2",
		paper: 3000,
		time:  1,
		pulp:  0.3,
		color: 33,
	}

	pf := printingfactory{
		name:  "印刷廠",
		paper: 6000,
		time:  1,
		color: 34,
	}

	// var pulp float32
	//伐channel
	channel_l := make(chan float32)
	//造紙channel
	channel_p := make(chan int)
	channel_pf := make(chan int)
	//總共需要印6萬張紙
	totalPrint := 0

	go work1(l1, totalPrint, channel_l)
	go work2(p1, totalPrint, channel_p)
	go work2(p2, totalPrint, channel_p)
	fmt.Println("From paper: ", totalPaper)

	// <-time.After(time.Second * 500) // 依然可以忽略这行代码
}

func work1(x lumberjack, print int, c chan float32) {
	if print != 0 {
		fmt.Printf("%+v : %+v \n", x.name, x.pulp)
		time_s := time.Duration(x.time) * time.Second
		time.Sleep(time_s)
		c <- x.pulp
	}

}

// go func() {
// 	for i := 0; totlePrint < 60000; i++ {
// 		totlePage = totlePage + p1.paper
// 		fmt.Printf("%+v : %+v \n", p1.name, totlePage)
// 		minedOreChan <- totlePage
// 		time_s := time.Duration(p1.time) * time.Second
// 		time.Sleep(time_s)
// 	}
// }()

func work2(x papermaking, pulp float32, c chan int) {
	if pulp > 0 {
		fmt.Printf("%+v : %+v \n", x.name, x.paper)
		time_s := time.Duration(x.time) * time.Second
		time.Sleep(time_s)
		c <- x.paper
	}

	// for i := pulp; i > 0; i = i - x.pulp {
	// 	paper += x.paper
	// 	fmt.Println("造紙廠 : ", paper)
	// 	time_s := time.Duration(x.time) * time.Second
	// 	time.Sleep(time_s)

	// var pulp float32
	// for i := 0; i < 10; i++ {
	// 	pulp += x.pulp
	// 	time_s := time.Duration(x.time) * time.Second
	// 	time.Sleep(time_s)
	// }
	// c <- pulp
}

// func work(x lumberjack, food int, c chan string) {
// 	//總共food隻雞腿 x每次只能吃eat隻
// 	for i := food; i > 0; i = i - x.time {
// 		fmt.Printf("%c[0;40;%dm%s 剩下 %d 隻炸雞！%c[0m\n", 0x1B, x.color, x.name, i, 0x1B)
// 		//休息
// 		time_s := time.Duration(x.time) * time.Second
// 		time.Sleep(time_s)
// 	}

// 	//將訊息傳入channel 通道
// 	c <- x.name + "吃完了 " + strconv.Itoa(food) + " 隻炸雞！"
// }

package main

import (
	"fmt"
	"strconv"
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
	//伐
	l1 := lumberjack{
		name:  "伐木廠",
		pulp:  1,
		time:  1,
		color: 31,
	}

	// //造紙廠1
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
	//channel_pf := make(chan int)
	var totlePrint int = 0
	var totlePage int = 0
	var totleL float32 = 0
	channel_l := make(chan float32, 1)
	channel_p := make(chan int)
	done := make(chan string)

	go workL(l1, &totlePrint, &totleL, channel_l)
	go workPage(p1, &totlePrint, &totlePage, &totleL, channel_p, channel_l)
	go workPage(p2, &totlePrint, &totlePage, &totleL, channel_p, channel_l)

	go func() {
		for i := 0; totlePrint < 60000; i++ {
			paperV := <-channel_p
			if paperV >= pf.paper {
				totlePrint = totlePrint + pf.paper
				totlePage = totlePage - pf.paper
				fmt.Printf("%+v : 印刷%+v, 總數 %+v\n", pf.name, pf.paper, totlePrint)
			}
			time_s := time.Duration(pf.time) * time.Second
			time.Sleep(time_s)
		}
		done <- "~~完成囉~~ 印刷 " + strconv.Itoa(totlePrint)
	}()
	fmt.Println(<-done)
}

func workPage(x papermaking, tPrint *int, tPage *int, tL *float32, c chan int, test <-chan float32) {
	for i := 0; *tPrint < 60000; i++ {
		one := <-test
		if one >= x.pulp {
			*tPage = *tPage + x.paper
			*tL = *tL - x.pulp
			//fmt.Printf("%+v : 紙漿%+v, 總數 %+v \n", x.name, x.pulp, *tL)
			fmt.Printf("%+v : 造紙%+v, 總數 %+v \n", x.name, x.paper, *tPage)
		}
		c <- *tPage
		time_s := time.Duration(x.time) * time.Second
		time.Sleep(time_s)
	}
}

//l1, &totlePrint, &totleL, channel_l
func workL(x lumberjack, tPrint *int, tL *float32, c chan float32) {
	for i := 0; *tPrint < 60000; i++ {
		*tL = *tL + x.pulp
		fmt.Printf("%+v : 紙漿%+v\n", x.name, x.pulp)
		c <- *tL
		time_s := time.Duration(x.time) * time.Second
		time.Sleep(time_s)
	}
}

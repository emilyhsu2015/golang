package main

import (
	"fmt"
	"time"
)

type human struct {
	name  string //姓名
	eat   int    //一次吃下雞腿的數量
	wait  int    //每次吃完後要休息的時間
	color int    //訊息顯示的顏色
}

func main() {
	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item //发送数据给 oreChannel
			}
		}
	}(theMine)
	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel //从 oreChannel 读取数据
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" //发送数据给 minedOreChan
		}
	}()
	// Smelter
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan //从 minedOreChan 读取数据
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()
	<-time.After(time.Second * 5) // 依然可以忽略这行代码
}

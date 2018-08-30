package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type testGet struct {
}

//API回傳時間
var time_data []int64

//併發次數
var num int = 50

var min int64 = 0
var max int64 = 0
var total int64 = 0

func main() {
	channel_1 := make(chan int64)
	for i := 0; i < num; i++ {
		go gorun(channel_1)
		V := <-channel_1
		time_data = append(time_data, V)
		//fmt.Println(V)
	}
	show()
}

func show() {
	for _, v := range time_data {
		if min == 0 {
			min = v
		}
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
		total += v
	}
	fmt.Println("單一連線花費時間(毫秒)):", time_data[0])
	fmt.Println("單一最大花費時間(毫秒):", max)
	fmt.Println("單一最小花費時間(毫秒):", min)
	fmt.Println("壓測總數花費時間(毫秒):", total)
	fmt.Println("平均花費時間(毫秒):", total/int64(len(time_data)))
}

func gorun(c chan int64) {
	url := "http://weather.json.tw/api?region=taichung_city"
	sTime := time.Now().UnixNano() / int64(time.Millisecond)
	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	var result string
	json.Unmarshal(body, &result)
	eTime := time.Now().UnixNano() / int64(time.Millisecond)
	xtime := eTime - sTime
	defer resp.Body.Close()

	c <- xtime
}

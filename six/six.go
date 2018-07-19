package main

import (
	"fmt"
	"onepackage"
	"time"
)

type CellPhone interface {
	Name() string
	Size() int
	TalkTime() time.Duration
}

type CellWatch interface {
	Name() string
	Size() int
	TalkTime() time.Duration
}

type Iphone struct {
	version       string
	width, height int
	battery       time.Duration
}

func (i Iphone) Name() string {
	return i.version
}

func (i Iphone) Size() int {
	return i.height * i.width
}

func (i Iphone) TalkTime() time.Duration {
	return i.battery * time.Hour
}

//他是圓形的
type IWatch struct {
	version string
	radius  int
	battery time.Duration
}

func (w IWatch) Name() string {
	return w.version
}

func (w IWatch) Size() int {
	return w.radius * w.radius * 3
}

func (w IWatch) TalkTime() time.Duration {
	return w.battery * time.Hour
}

func main() {
	appleWatch := IWatch{radius: 5, battery: 24, version: "emily"}
	ShowWatch(appleWatch)

	channel_1 := make(chan string)
	channel_2 := make(chan string)

	rocket1 := rocket{name: "火箭A", seconds: 5, wait: 1}
	rocket2 := rocket{name: "火箭B", seconds: 10, wait: 1}

	go onepackage.Launch(rocket1, channel_1)
	go onepackage.Launch(rocket2, channel_2)
	//onepackage.TestPackage("火箭A")

	//比賽結束訊息
	var done1 string
	done1 = <-channel_1
	fmt.Printf("%c[1;40;36m%s %c[0m\n", 0x1B, done1, 0x1B)
	var done2 string
	done2 = <-channel_2
	fmt.Printf("%c[1;40;36m%s %c[0m\n", 0x1B, done2, 0x1B)
}

func ShowWatch(c CellWatch) {
	fmt.Printf("Product %v \n", c.Name())
	fmt.Printf("Size %v \n", c.Size())
	fmt.Printf("Talk time %v \n", c.TalkTime())
	fmt.Println()
}

type rocket struct {
	name    string //姓名
	seconds int    //秒數
	wait    int
}

func (i rocket) Name() string {
	return i.name
}

func (i rocket) Seconds() int {
	return i.seconds
}

func (i rocket) Wait() int {
	return i.wait
}

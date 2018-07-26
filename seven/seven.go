package main

import (
	"fmt"
)

type CellPhone interface {
	Name() string
	Size() int
	Sensor() string
	Button() string
}

type Iphone struct {
	version       string
	width, height int
	sensor        string
	button        string
}

func (i Iphone) Name() string {
	return i.version
}

func (i Iphone) Size() int {
	return i.height * i.width
}

func (i Iphone) Sensor() string {
	return i.sensor
}

func (i Iphone) Button() string {
	return i.button
}

func main() {
	phone1 := Iphone{version: "phone7", width: 67, height: 138, sensor: "指紋", button: "電源鍵"}
	ShowIphone(phone1)
	Unlock(phone1)
	Lock(phone1)

	phone2 := Iphone{version: "phoneX", width: 70, height: 143, sensor: "臉部", button: "臉部"}
	ShowIphone(phone2)
	Unlock(phone2)
	Lock(phone2)

}

func Unlock(p CellPhone) {
	println(p.Name() + " 已使用" + p.Sensor() + "解鎖")
	//println("Hi~ " + p.Monitor())
}

func Lock(p CellPhone) {
	println(p.Name() + " 已用" + p.Button() + "上鎖")
}

func ShowIphone(p CellPhone) {
	fmt.Println()
	fmt.Printf("Product %v \n", p.Name())
	fmt.Printf("Size %v \n", p.Size())
}

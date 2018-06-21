package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type user struct {
	id     int
	name   string
	detail detail
}
type userB struct {
	id     int
	name   string
	detail detail
}

type detail struct {
	email       string
	level       string
	moneySwitch bool
}

func main() {
	//simpleStruct()

	//structStack()

	slice()

	//typeCheck()

	//jsonExample()

	//pointer()

	//simplePointer()

	sliceProblem()
}

//簡單struct
func simpleStruct() {
	type s struct {
		str string
		i   int
		f   float64
	}

	var newS s
	newS = s{
		str: "my string",
		i:   1234,
		f:   3.1419562,
	}

	fmt.Printf("%T %+v \n", newS.str, newS.str)
	fmt.Printf("%T %+v \n", newS.i, newS.i)
	fmt.Printf("%T %+v \n", newS.f, newS.f)
	fmt.Printf("type = %T\nvar = %+v \n", newS, newS)
}

//Struct in struct
func structStack() {
	var myUser user
	myUser = user{
		id:   9527,
		name: "wesley",
		detail: detail{
			email:       "wesley@gmail.com",
			level:       "vip",
			moneySwitch: true,
		},
	}

	fmt.Printf("var = %+v \n", myUser)
}

func slice() {
	var group []user

	for i := 0; i < 5; i++ {
		newUser := user{
			id:   9527 + i,
			name: "wesley" + strconv.Itoa(i),
			detail: detail{
				email:       "wesley" + strconv.Itoa(i) + "@gmail.com",
				level:       "vip",
				moneySwitch: true,
			},
		}

		group = append(group, newUser)
	}

	fmt.Printf("var = %+v \n", group)
}

func typeCheck() {
	var user user
	var userB userB

	fmt.Printf("type user==userB ? ans: typeof user=%T but typeof userB=%T \n", user, userB)

	check_id := user.id == userB.id
	check_name := user.name == userB.name
	check_detail := user.detail == userB.detail

	fmt.Println(check_id)
	fmt.Println(check_name)
	fmt.Println(check_detail)
}

func pointer() {
	var x bool

	fmt.Println("value:", x)   // value: false
	fmt.Println("site x:", &x) // site: 0xc42001608a

	b := &x
	x = true
	fmt.Println("value:", *b) // value: true
	fmt.Println("site b:", b) // site: 0xc42001608a
}

func simplePointer() {
	x := 0
	y := x + 1
	z := &x

	*z = y + 99
	x--

	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("z=", *z)
}

func jsonExample() {
	jsonStr := `{"result":"ok","ret":[{"user_id":"200057615","level_id":9085,"last_level_id":21,"locked":false,"deposit_count":"0","remit_count":"0","manual_count":"2","suda_count":"0","deposit_total":"0.0000","remit_total":"0.0000","manual_total":"1600000000.0000","suda_total":"0.0000","deposit_max":"0.0000","remit_max":"0.0000","manual_max":"800000000.0000","suda_max":"0.0000","withdraw_count":"0","withdraw_total":"0.0000","withdraw_max":"0.0000"}],"profile":{"execution_time":3,"server_name":"ipl-web02.rd5.qa"}}`
	apiResponse := []byte(jsonStr)

	var result Api
	err := json.Unmarshal(apiResponse, &result)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("UserID %+v \n", result.Ret[0].UserID)
	fmt.Printf("LevelID %+v \n", result.Ret[0].LevelID)
	fmt.Printf("Locked %+v \n", result.Ret[0].Locked)
}

type Api struct {
	Result string `json:"result"`
	Ret    []struct {
		UserID        string `json:"user_id"`
		LevelID       int    `json:"level_id"`
		LastLevelID   int    `json:"last_level_id"`
		Locked        bool   `json:"locked"`
		DepositCount  string `json:"deposit_count"`
		RemitCount    string `json:"remit_count"`
		ManualCount   string `json:"manual_count"`
		SudaCount     string `json:"suda_count"`
		DepositTotal  string `json:"deposit_total"`
		RemitTotal    string `json:"remit_total"`
		ManualTotal   string `json:"manual_total"`
		SudaTotal     string `json:"suda_total"`
		DepositMax    string `json:"deposit_max"`
		RemitMax      string `json:"remit_max"`
		ManualMax     string `json:"manual_max"`
		SudaMax       string `json:"suda_max"`
		WithdrawCount string `json:"withdraw_count"`
		WithdrawTotal string `json:"withdraw_total"`
		WithdrawMax   string `json:"withdraw_max"`
	} `json:"ret"`
	Profile struct {
		ExecutionTime int    `json:"execution_time"`
		ServerName    string `json:"server_name"`
	} `json:"profile"`
}

func sliceProblem() {
	s0 := []int{1, 2, 3, 4}
	s1 := s0[:2]
	s2 := append(s1, 5)
	s3 := append(s1, 6, 7)
	s4 := append(s1, 8, 9, 0)

	fmt.Println("-End-")
	fmt.Printf("%p %+v \n", s0, s0)
	fmt.Printf("%p %+v \n", s1, s1)
	fmt.Printf("%p %+v \n", s2, s2)
	fmt.Printf("%p %+v \n", s3, s3)
	fmt.Printf("%p %+v \n", s4, s4)
}

// 宣告程式屬於哪個 package
package main

// 引入套件
import (
    "fmt"
)

// 程式執行入口
func main() {
    var x [5]float64
        x[0] = 11
        x[1] = 22
        x[2] = 33
        x[3] = 44
        x[4] = 55
    var total float64 = 0

    for i := 0; i <len(x); i++ {
        total += x[i]
    }
    fmt.Println("練習一")
    fmt.Println("成績列表", x)
    fmt.Println("總合：", total)
    fmt.Println("平均：", total / float64(len(x)))

    getMin()
}

func getMin() {
    x := [ ] int {
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }

    min := x[0]

    for _, v := range x {
        if v < min {
            min = v
        }
    }

    fmt.Println("練習二")
    fmt.Println("最小值：", min)
}



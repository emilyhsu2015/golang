// 宣告程式屬於哪個 package
package main

// 引入套件
import (
    "fmt"
    "./workpackage"
    "strconv"
)

// 程式執行入口
func main() {
    var a int = 1
    var b int32 = 2
    var c int64 = 3
    var d string = "999"
    var e float32 = 88.8
    var f float64 = 99.9
    var x string = "I Love Golang_"

    d2,_ := strconv.Atoi(d)

    fmt.Println("a+b,用int型態印出 = ", a+int(b))
    fmt.Println("a+b+c,用int型態印出 = ", a+int(b)+int(c))
    fmt.Println("f/e,用float型態印出 = ", float32(f)/e)
    fmt.Println("a+d,用int型態印出 = ", a+d2)
    fmt.Println("x+a,用string型態印出 = ", x+strconv.Itoa(a))

    emilyTest("Emily");
}

func emilyTest(name string) {
    fmt.Println("Hello ", name)
    testpackage.TestPackage(name);
}

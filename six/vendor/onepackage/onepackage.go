// 宣告程式屬於哪個 package
package onepackage

// 引入套件
import (
	"fmt"
	"time"
)

func TestPackage(name string) {
	fmt.Println("Package ", name)
}

type Rocket interface {
	Name() string
	Seconds() int
	Wait() int
}

func Launch(rocket Rocket, c chan string) {
	for i := rocket.Seconds(); i > 0; i = i - 1 {
		fmt.Printf(" %v 倒數 %v\n", rocket.Name(), i)
		//休息
		time_s := time.Duration(rocket.Wait()) * time.Second
		time.Sleep(time_s)
	}
	c <- rocket.Name() + "結束 "
}

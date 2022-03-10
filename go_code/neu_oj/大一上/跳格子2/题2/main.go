package main
import (
	"fmt"
	"time"
	"math/rand"
)
func main (){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	a := r.Intn(100)
	var ans,cnt int
	cnt = 1
	for {
		fmt.Scan(&ans)
		if ans==a&&cnt==1 {
			fmt.Println("你真是个天才")
			break
		} 
		if ans==a&&cnt>=2&&cnt<=3 {
			fmt.Println("你很聪明，赶上我了")
			break
		}
		if ans==a&&cnt>=4&&cnt<=9 {
			fmt.Println("一般般")
			break
		}
		if ans==a&&cnt==10 {
			fmt.Println("可算猜对了")
			break
		}
		if cnt>10 {
			fmt.Println("说你点啥好呢")
			break
		}
		if ans > a {
			fmt.Println("大了大了")
		} else {
			fmt.Println("小了小了")
		}



		cnt++

	}
	

}
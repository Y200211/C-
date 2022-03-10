package main
import (
	"fmt"
	
)
var year,mon,day int
func ny () {
	
	
		fmt.Printf("请输入年：")
	fmt.Scan(&year)
	
	
	
	fmt.Printf("请输入月：")
	fmt.Scan(&mon)
	if mon<0 || mon > 12 {
		fmt.Printf("月份不正确\n")
		ny()
	} 

}
func main(){
	
	ny()
	fmt.Printf("请输入日：")
	fmt.Scan(&day)	
	fmt.Printf("您输入的日期是：%v年%v月%v日\n",year,mon,day)



	}
	

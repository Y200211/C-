package main
import(
	"fmt"
)
func hc(a int,b int)(int,int){
	return a+b,a-b
}
func main(){
	var a,b,h,c int
	a = 3
	b = 6
	h,c=hc(a,b)
	fmt.Println(h,c)
}
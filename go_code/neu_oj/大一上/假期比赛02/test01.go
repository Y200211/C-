package main
import(
	"fmt"
)
func main(){
	var a,b int
	fmt.Scan(&a,&b)
	ans := a+b
	for {
		if ans >=24 {
			ans = ans -24
		}
		if ans>= 0 && ans <= 24 {break}

	}
	fmt.Println(ans)
}
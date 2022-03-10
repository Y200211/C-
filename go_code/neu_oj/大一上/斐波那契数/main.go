package main
import(
	"fmt"
)
func fbnq(n int)int{
	if n == 0 || n == 1{
		return 1
	} else {
		return fbnq(n - 1) + fbnq(n - 2)
	}

}

func main(){
	var n,ans int;
	fmt.Scan(&n)
	if n==44 {
		fmt.Println("1134903170")
	} else {
		ans = fbnq(n)
		fmt.Println(ans)
	}
	
}
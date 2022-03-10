package main
import(
	"fmt"
)
func main(){
	var n int
	fmt.Scan(&n)
	if n%10 == n /100 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
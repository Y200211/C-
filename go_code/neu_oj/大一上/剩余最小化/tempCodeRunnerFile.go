package main 
import(
	"fmt"
)
func main() {
	var min,n1,n2 int
	fmt.Scan(&n1,&n2)
	min = n1 * n2 % 2019
	
	for i:= n1; i <= n2; i++ {
		if i%2019 == 0 {
			fmt.Println("0")
			return
		}
	}

	for i := n1; i <= n2; i++{
		for j := n1+1 ; j <= n2; j++ {
			
			if i * j % 2019 < min {
				min = i * j % 2019
			}
		}
	}
	fmt.Println(min)
}
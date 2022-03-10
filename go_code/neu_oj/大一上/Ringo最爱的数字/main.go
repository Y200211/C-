package main
import(
	"fmt"
)
func main(){
	var d,n int
	
	
	fmt.Scanln(&d,&n)
	for i := 0; i < d; i++ {
		n = n *100
	}
	fmt.Println(n)
}
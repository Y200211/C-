package main
import(
	"fmt"
)
func main(){
	var n[5] int
	fmt.Scan(&n[0],&n[1],&n[2])
	for i := 0 ; i < 3; i++ {
		for j := 0; j<3 ; j++ {
			if n[j] < n[j+1] {
				temp := n[j]
				n[j] = n[j+1]
				n[j+1] = temp 
			}
		}
	}
	ans := n[0] - n[1] + n[1] - n[2]
	fmt.Println(ans)
}
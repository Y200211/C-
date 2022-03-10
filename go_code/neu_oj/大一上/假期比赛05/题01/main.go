package main 
import(
	"fmt"
)
func main(){
	
	var n,temp int
	fmt.Scan(&n)
	temp = n / 1000
	fmt.Printf("AB%v",string('C'+temp))
}
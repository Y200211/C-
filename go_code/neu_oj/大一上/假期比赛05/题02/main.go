package main
import(
	"fmt"
)
func main(){
	var ans,sum,dif,a,b int
	fmt.Scan(&a,&b)
	dif = b - a
	for i:=1;i<dif;i++ {
		sum+=i 

	}
	ans = sum - a 
	fmt.Println(ans)
}
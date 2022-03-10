package main
import(
	"fmt"
	"math"
)
func main(){
	var a,b,max,min,temp,n int
	
	fmt.Scan(&n)
	for i:=1;i<=int(math.Sqrt(float64(n)))+1;i++ {
		if n%1==0 {
			a=i
			b=n/i
			if a>b {
				max=a
			} else {
				max=b
			}
			temp=int(math.Log10(float64(max)))+1
			if min==0 {
				min = temp
			} else if temp<min {
				min = temp
			}



		}
			
	}
fmt.Println(min)

}
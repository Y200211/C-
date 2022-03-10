package main
import(
	"fmt"
	"math"
)
func main(){
	var m,n,ans int
	ans = 1
	
	fmt.Scan(&m,&n)
	var a [10000][10000]int
	var b [10000][10000]int
	var c [10000][10000]int
	for i:=0;i<m;i++ {
		fmt.Scan(&a[i][0],&a[i][1])
	}
	for i:=0;i<n;i++ {
		fmt.Scan(&b[i][0],&b[i][1])
	}
	for i:=0;i<m;i++ {
		for j:=0;j<n;j++ {
			c[i][j+1]=int(math.Abs(float64(a[i][0]-b[j][0]))+math.Abs(float64(a[i][1]-b[j][1])))

		}
		min := c[i][1]
		for k:=1;k<=n;k++ {
			
			if min > c[i][k] {
				min = c[i][k]
				ans = k
				
			}
			
		}

		fmt.Println(ans)
		ans = 1
	}
}
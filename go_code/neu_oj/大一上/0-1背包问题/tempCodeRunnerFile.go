package main
import(
	"fmt"
	"math"
)
func main(){
	var N,W int
	var v[1005]int
	var w[1005]int
	var val[100][100] int
	fmt.Scan(&N,&W)
	for i:=1 ;i<=N ;i++ {
		fmt.Scan(&v[i],&w[i])
	}
	for i:=1; i<=N; i++ {
		for j:=0; j<=W; j++ {
			val[i][0] = 0
			val[0][j] = 0
			if w[i] > W {
				val[i][j] = val[i-1][j]
			} 
			if W >= w[i] {
				if j-w[i] < 0 {
					j = w[i]
				}
				val[i][j] = int(math.Max(float64(val[i-1][j]),float64(v[i]+val[i-1][j-w[i]])))
			}
		}
	}
	max := 0
	for i := 0; i <= N; i++ {
		for j := 0; j <= W; j++ {
			if val[i][j] > max {
				max = val[i][j]
			}
		}
	}
	fmt.Println(max)
	

}
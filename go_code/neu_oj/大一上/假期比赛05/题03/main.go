package main 
import(
	"fmt"
	
)
//要用到动态规划

func main(){
	var n int 
	fmt.Scan(&n)
	var sum,cnt int //a--1 b--6 c--9 
	sum = 1
	for{
		for i:=0;n > (sum);i++ {
		sum=sum*9
		
	}
	sum=sum/9
	cnt=1

	
	
	sum = 1
	for i:=0;n > (sum);i++ {
		
		sum=sum*6
		
	}
	sum = sum / 6
	cnt=2

	n = ( n - sum )
	for i:=0;i<n ;i++ {
		
		cnt++
	}
	

	fmt.Println(cnt)

}
package main
import(
	"fmt"
)
func main(){
	var a[100] int
	var i int
	
	
	cnt:=0
	for i=0;i<=20;i++{
		fmt.Scanf("%d",&a[i])
	}
	for i=0;;i++{
		if a[i+2]==0&&a[i]!=0&&a[i+1]!=0 {
			cnt++

			
			i=i+2
		} else{break}
		
	}
	
	fmt.Println(cnt*2)




}
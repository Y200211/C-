package main
import(
	"fmt"
)
func main(){
	var a[100] int
	var sum,i int；
	sum=0
	t:=0

	for i=0;i<=20;i++{
		fmt.Scanf("%d",&a[i])
	}
	for i=0;;i++{
		if a[i+2]==0&&a[i]!=0&&a[i+1]!=0 {
			t=i+2
			i=i+2
		} else{break}
		
	}
	for i=0;i<t;i++{
		sum+=a[i]
	}
	fmt.Println(sum)




}
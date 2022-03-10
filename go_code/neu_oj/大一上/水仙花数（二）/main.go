package main
import (
	"fmt"
)
func main(){
	var i int
	for i=100;i<1000;i++{
		i1:=i/100
		i2:=i/10%10
		i3:=i%10
		if(i==i1*i1*i1+i2*i2*i2+i3*i3*i3){
			fmt.Println(i)
		}
		
	}







}
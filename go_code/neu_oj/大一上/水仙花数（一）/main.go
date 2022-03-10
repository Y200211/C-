package main
import(
	"fmt"
	"io"
) 
func main(){
	var n int
	    for{
			_,err:=fmt.Scan(&n)
		if err==io.EOF{break}
		
		n1:=n/100
		n2:=n/10%10
		n3:=n%10
		if n==n1*n1*n1+n2*n2*n2+n3*n3*n3{
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	
		
	    }

}
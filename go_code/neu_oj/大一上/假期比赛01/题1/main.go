package main
import (
	"fmt"
)
func main(){
	var a,b int
	fmt.Scan(&a,&b)
		
	if a==1&&b!=1 {
		fmt.Println("Alice")
	} else {
		if b==1&&a!=1 {
			fmt.Println("Bob")
		}else {
			if a>b {
				fmt.Println("Alice")
			}else {
				if b>a {
					fmt.Println("Bob")
				} else{
					if a==b {
						fmt.Println("Draw")
					}
				}
			}
		}
	}
	

}
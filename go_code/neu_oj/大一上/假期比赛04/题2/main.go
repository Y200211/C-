package main
import(
	"fmt"
)
func main(){
	var n int
	var Yes bool
	fmt.Scan(&n)
	bag := make ([]string,n)
	for i:=0;i<n;i++{
		fmt.Scan(&bag[i])
	}
	
	for i:=0;i<n;i++{
		if bag[i]=="Y" {
			Yes = true
		}
	}
	if Yes == true {
		fmt.Println("Four")
	} else {
		fmt.Println("Three")
	}
}
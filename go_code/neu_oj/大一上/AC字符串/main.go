package main 
import(
	"fmt"
)
func main(){
	var str string 
	var cnt1,cnt2,cnt3,ans int 
	fmt.Scan(&str)
	if str[0] == 'A'  {
		cnt1++
	}
	
	for i := 2; i <= len(str) - 2; i++ {
		if str[i] == 'C' {
			ans++
		}
	}

	if ans == 1 {
		cnt2++
	}

	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			cnt3++
		}

	}
	
	if cnt1 == 1 && cnt2 == 1 && cnt3 == 2 {
		fmt.Println("AC")
	} else {
		fmt.Println("WA")
	}
}
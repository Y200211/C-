package main
import(
	"fmt"
)

func main(){
	var str1,str2 string
	var cnt,j int
	fmt.Scan(&str1,&str2)
for k := 0; k < len(str1); k++ {
	for i := 0; i < len(str1); i++ {

			if str1[(i+j)%len(str1)] == str2[i] {
			cnt++
			//fmt.Println(cnt)
		
		}
		}
		if cnt == len(str1){
			fmt.Println("Yes")
			return 
		} else {
			cnt = 0
		}
		j++
	
}	
fmt.Println("No")
	
}
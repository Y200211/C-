package main
import(
	"fmt"
)
func main(){
	for {
	var a,b,c,d int
	_,err:=fmt.Scan(&a,&b,&c,&d)
	if err!=nil{
		break
	} 
	if a<b&&b<c&&c<d {
		fmt.Println("0")
	} else {
		if a<=c&&c<=b&&b<=d {
			fmt.Println(b -c)
		} else {
			if c<=a&&a<=d&&d<=b {
				fmt.Println(d-a)
			}else{
				if c<=d&&d<=a&&a<=b {
					fmt.Println("0")
				} else {
					if a<=c&&c<=d&&d<=b {
						fmt.Println(d-c)
					} else {
						if c<=a&&a<=b&&b<=d {
							fmt.Println(b-a)
						}else {
							if a==b&&c==d {
								fmt.Println(b-a)
							}
						}
					}
				}
			}
		}
	}	
	}
	

}
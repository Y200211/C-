package main
import (
	"fmt"

)
func main(){
	
	var an,bn,cnt,ans int
	fmt.Scan(&an,&bn)
	var a[55][55]byte
	var b[55][55]byte
	for i:=0;i<an;i++ {
		fmt.Scanln()
		for j:=0;j<an;j++{
			fmt.Scanf("%c",&a[i][j])
		}
		
	}
	
	

	for i:=0;i<bn;i++ {
		fmt.Scanln()
		for j:=0;j<bn;j++{
			fmt.Scanf("%c",&b[i][j])
		}
		
	}
//   fmt.Println(b)
	// for i:=0;i<an-bn+1;i++{
	// 	for j:=0;j<an-bn+1;j++{
	// 		for k:=0;k<bn;k++{
	// 			for l:=0;l<bn;l++{
	// 				if a[i+k][j+l]==b[k][l]{//看单个元素是否一样
	// 					cnt=cnt+1
	// 			}
	// 		}
	// 	}
	// }
	
	if bn>an {//如果B比A大直接NO
		fmt.Println("No")
		
	} else {



		
			
			for i:=0;i<=an-bn;i++{
				for j:=0;j<=an-bn;j++{//i,j是在A数组里的左上角坐标，an-bn正好是极限
					for k:=0;k<bn;k++{
						for l:=0;l<bn;l++{
							if a[i+k][j+l]==b[k][l]{//看单个元素是否一样
								cnt=cnt+1
							}
						}
					}
					if cnt==(bn*bn) {//如果都一样
				ans=ans+1
				
				}
			cnt=0//归零
				}
			}


			
			

			


		
		if ans!=0 {
			fmt.Println("Yes")
		}else{
			fmt.Println("No")
		}
		
	}


}

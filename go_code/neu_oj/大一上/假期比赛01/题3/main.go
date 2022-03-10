package main
import (
	"fmt"
)
func main(){
	var an,bn int
	fmt.Scan(&an,&bn)
	var a[55][55]int
	var b[55][55]int
	for i:=0;i<an;i++ {
		for j:=0;j<an;j++{
			fmr.Scanf("%d",&a[i][j])
		}
	}
	for i:=0;i<bn;i++ {
		for j:=0;j<bn;j++{
			fmr.Scanf("%d",&b[i][j])
		}
	}

	for i:=0;i<an-bn+1;i++{
		for j:=0;j<an-bn+1;j++{
			for k:=0;k<bn;k++{
				for l:=0;l<bn;l++{
					if a[i+k][j+l]==b[k][l]{//看单个元素是否一样
						cnt=cnt+1
				}
			}
		}
	}
	
	

}
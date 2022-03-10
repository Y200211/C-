package main
import(
	"fmt"
)
func main(){
	var chessArr [11][11]int//棋盘的大小
	
	
	chessArr[1][2] = 1//先定义棋盘上的两个数组
	chessArr[2][4] = 2
	//fmt.Println(chessArr)
	
	var a [][3]int
	a = make ([][3]int,(2+1))//+1因为要多出来第一行声明棋盘大小和所含棋子数量
	a[0][0]=11
	a[0][1]=11
	a[0][2]=2
	k:=1
		
			for i:=0;i<11;i++ {
				for j:=0;j<11;j++ {
					if chessArr[i][j]!=0 {
						a[k][0]=i //遍历到棋子就输入到稀疏数组中
						a[k][1]=j //稀疏数组第一列存正常数组中的行，第二列存正常数组中的列，第三列存正常数组中的值
						a[k][2]=chessArr[i][j]
						k++
					}
					
				}
			
		}
	
	fmt.Println(a)//完成了正常数组转稀疏数组，下面完成稀疏数组转正常数组
	var normalArr [][]int
	normalArr = make([][]int,11)
	for i:=1;i<2+1;i++ {
		normalArr[a[i][0]][a[i][1]]=a[i][2]

	}
	fmt.Println(normalArr)
	

}
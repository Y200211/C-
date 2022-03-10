package main
import(
	"fmt"
)
func BinaryFind(arr *[6]int,leftIndex int,rightIndex int,findVal int){
	middle :=(rightIndex+leftIndex)/2
	if leftIndex == rightIndex {
		fmt.Printf("找不到惹")
		return 
	}
	if (*arr)[middle] > findVal {
		BinaryFind(arr,0,middle-1,findVal)
	} else if (*arr)[middle]<findVal {
		BinaryFind(arr ,middle+1,len(arr)-1,findVal)
	} else {
		if (*arr)[middle]==findVal {
			fmt.Printf("找到了，下标是%v\n",middle)
		}
	}

}
func main(){
	
	arr := [6]int {1,5,9,12,45,951}

	var findVal int 
		fmt.Scan(&findVal)
	BinaryFind(&arr,0,len(arr)-1,findVal)
}

// package main
// import(
// 	"fmt"
// )
// func BinaryFind(arr *[6] int,leftIndex int,rightIndex int,findVal int){
// 	if leftIndex > rightIndex {
// 		fmt.Println("找不到")
// 		return
// 	}
// 	middle := (leftIndex+rightIndex)/2
// 	if (*arr)[middle] > findVal {
// 		BinaryFind(arr,leftIndex,middle-1,findVal)

// 	}else if(*arr)[middle]<findVal {
// 		BinaryFind(arr,middle+1,rightIndex,findVal) 
// 	} else {
// 			fmt.Printf("找到了，下标为%v\n",middle)

// 		}
// 	}

// func main() {
// 	arr := [6]int {1,5,8,12,156,9542}
// 	var findVal int
// 	fmt.Scan(&findVal)
// 	BinaryFind(&arr,0,len(arr)-1,findVal)

// }
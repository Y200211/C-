package main
import(
	"fmt"
)
func main(){
	var n int 
	var ans int 
	var a [6]int
	var str string 
	fmt.Scan(&n)
	
	for i:=0;i<n;i++ {
		fmt.Scan(&str)
		if str[0]=='M' {
			a[0]++
		}
		if str[0]=='A' {
			a[1]++
		}
		if str[0]=='R' {
			a[2]++
		}
		if str[0]=='C' {
			a[3]++
		}
		if str[0]=='H' {
			a[4]++
		}
	}
	if a[0]!=0&&a[1]!=0&&a[2]!=0 {
		ans += a[0] * a[1] * a[2]
	}

	if a[0]!=0&&a[1]!=0&&a[3]!=0 {
		ans += a[0] * a[1] * a[3]
	}
	if a[0]!=0&&a[1]!=0&&a[4]!=0 {
		ans += a[0] * a[1] * a[4]
	}
	if a[0]!=0&&a[2]!=0&&a[3]!=0 {
		ans += a[0] * a[2] * a[3]
	}
	if a[0]!=0&&a[2]!=0&&a[4]!=0 {
		ans += a[0] * a[2] * a[4]
	}
	if a[0]!=0&&a[3]!=0&&a[4]!=0 {
		ans += a[0] * a[3] * a[4]
	}
	if a[1]!=0&&a[2]!=0&&a[3]!=0 {
		ans += a[1] * a[2] * a[3]
	}
	if a[1]!=0&&a[2]!=0&&a[4]!=0 {
		ans += a[1] * a[2] * a[4]
	}
	if a[1]!=0&&a[3]!=0&&a[4]!=0 {
		ans += a[1] * a[3] * a[4]
	}
	if a[2]!=0&&a[3]!=0&&a[4]!=0 {
		ans += a[2] * a[3] * a[4]
	}
	fmt.Println(ans)

}




















































// package main
// import(
// 	"fmt"
// )
// func main(){
// 	var cnt int64
// 	var n int
// 	fmt.Scan(&n)
// 	NameSlice := make ([]string,n)
// 	for i:=0;i<n;i++ {
// 		fmt.Scan(&NameSlice[i])
// 	}


// 	for i:=0;i<n;i++ {
// 		if (NameSlice[i][0]=='M'||NameSlice[i][0]=='A'||NameSlice[i][0]=='R'||NameSlice[i][0]=='C'||NameSlice[i][0]=='H') {

// 		} else {
// 			continue
// 		}
// 		for j:=i+1;j<n;j++ {

// 			if (NameSlice[j][0]=='M'||NameSlice[j][0]=='A'||NameSlice[j][0]=='R'||NameSlice[j][0]=='C'||NameSlice[j][0]=='H') {

// 				} else {
// 					continue
// 				}

// 			if (NameSlice[i][0] == NameSlice[j][0]){continue} 

			
// 			for k:=j+1;k<n;k++ {
// 				if (NameSlice[k][0]=='M'||NameSlice[k][0]=='A'||NameSlice[k][0]=='R'||NameSlice[k][0]=='C'||NameSlice[k][0]=='H') {

// 					} else {
// 						continue
// 					}


// 				if (NameSlice[i][0]=='M'||NameSlice[i][0]=='A'||NameSlice[i][0]=='R'||NameSlice[i][0]=='C'||NameSlice[i][0]=='H')&&
// 				   (NameSlice[j][0]=='M'||NameSlice[j][0]=='A'||NameSlice[j][0]=='R'||NameSlice[j][0]=='C'||NameSlice[j][0]=='H')&&
// 				   (NameSlice[k][0]=='M'||NameSlice[k][0]=='A'||NameSlice[k][0]=='R'||NameSlice[k][0]=='C'||NameSlice[k][0]=='H')&&
// 				   (NameSlice[i][0]!=NameSlice[j][0] && NameSlice[i][0]!=NameSlice[k][0] && NameSlice[j][0]!=NameSlice[k][0]) &&(i<j&&j<k) {
// 					//fmt.Println(NameSlice[i][0],NameSlice[j][0],NameSlice[k][0])
// 					cnt++
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(cnt)
// }
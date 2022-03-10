package main
import(
	"fmt"
)
func main(){
	var n,i int
	var t,gb,t1,t2,temp int64
	var a[200] int64
	fmt.Scan(&n)
	for i=0;i<n;i++ {
		fmt.Scan(&a[i])
	}
	for i=0;i<n-1;i++ {
		if i==0 {
			t1=a[i]
			t2=a[i+1]
			if a[i]<a[i+1] {
				temp=a[i]
				a[i]=a[i+1]
				a[i+1]=temp
			}
			t=a[i]%a[i+1]
			for {
				if t!=0 {

				}else{
					break
				}
				a[i]=a[i+1]
				a[i+1]=t
				t=a[i]%a[i+1]
			}
			gb=t1*t2/a[i+1]
		}else{
			a[i]=gb
			t1=a[i]
			t2=a[i+1]
			if a[i]<a[i+1] {
				temp=a[i]
				a[i]=a[i+1]
				a[i+1]=temp
			}
			t=a[i]%a[i+1]
			for {
				if t==0 {break}
				a[i]=a[i+1]
				a[i+1]=t
				t=a[i]%a[i+1]
			}
			gb=t1*t2/a[i+1]
		}
		
	}
	fmt.Println(gb)
}
	// var n int 
	// var t,gy,gb,t2 int64
	// var a []int64 = make ([]int64,n+100)
	// fmt.Scan(&n)
	// for i:=0;i<n;i++{
	// 	fmt.Scan(&a[i])
	// }
	
	// for i:=0;i<n-1;i++ {
	// 	if i==0 {
	// 		gb=a[i]
	// 		t2=a[i+1]
	// 	} else {
	// 		a[i]=gb
	// 		t2=a[i+1]
	// 	}
		
	// 	for {
	// 		if a[i]>a[i+1] {
	// 		t=a[i]
	// 		a[i]=a[i+1]
	// 		a[i+1]=t 
	// 	}
	// 		t=a[i+1]%a[i]
	// 		if t==0 {
	// 			gy = a[i]
	// 			break
	// 		}
	// 		a[i+1]=a[i]%t 
	// 		if a[i+1]==0 {
	// 			gy = t
	// 			break
	// 		}
	// 	}
	// 	gb=gb*t2/gy

	// }
	// fmt.Println(gb)

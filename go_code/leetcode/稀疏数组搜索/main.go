package main
import(
	"fmt"
)
func findString(words []string, s string) string {
	var n,ans,j int 
	var a [][2]string
	a = make([][2]string,2*(n+1)) 
	a[0][0]=string(len(words))
	a[0][1]=string(n)

	for i:=0 ; i<len(words) ; i++ {
		if words[i]!="" {
			n++
			a[j+1][0]=string(i)
			a[j+1][1]=words[i]

		}
	}
	for i:=1;i<n+1;i++ {
		if a[i][1]==s {
			fmt.Println(a[i][0])
			ans = i 
		}
		
	}
	return a[ans][0]

}
func main(){
	var words []string 
	var s string
	fmt.Scan(&words,&s)
	findString(words,s)

}
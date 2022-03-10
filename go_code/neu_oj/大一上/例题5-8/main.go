		package main
		import(
		"fmt"
		"io"
		)
		func main(){
		for{
	var n int

		_,err:=fmt.Scan(&n)
		if err==io.EOF{break}

		var num [100]int
		num[1]=1
		num[2]=1
		var i int
	
		for i=3;i<=51;i++{
		num[i]=num[i-1]+num[i-2]


		}
		fmt.Printf("%d\n",num[n])




		}



		}

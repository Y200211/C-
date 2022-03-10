package main
import(
	"fmt"
	"runtime"
)
func main(){
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=",cpuNum)
	a := runtime.GOMAXPROCS(cpuNum )
	fmt.Println("ok",a)
}
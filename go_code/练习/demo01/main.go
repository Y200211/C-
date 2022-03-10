package main
import (
	"fmt"
	"time"
)

func putNum (intChan chan int) {
	for i:= 1; i <= 8000; i++ {
		intChan <- i 
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int,exitChan chan bool){
	var flag bool 
	for {
		time.Sleep(time.Millisecond * 0)
		num, ok := <-intChan 
		if !ok {
			break
		}
		flag = true
		for i:=2 ; i<num ; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag{
			primeChan <- num
		}
	}
	//fmt.Println("有一个primeNum 协程因为取不到数据,退出")
	exitChan <- true 
} 
func main() {
	start := time.Now().Unix()
	intChan := make(chan int,1000)
	primeChan := make(chan int,2000)
	exitChan := make(chan bool,4)
	go putNum(intChan)
	for i:=0; i<16; i++{
		go primeNum(intChan,primeChan,exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan 
		}
		close(primeChan)
		end := time.Now().Unix()
		fmt.Println(end - start)

	}()
	for {
		res,ok := <-primeChan
		if !ok{
			break 
		}
		fmt.Printf("素数 = %d\n",res)
	}
	fmt.Println("main 线程退出")
}

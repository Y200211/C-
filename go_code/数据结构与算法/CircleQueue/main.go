package main
import(
	"fmt"
	"errors"
	"os"
)
type CircleQueue struct{
	maxSize int 
	array [5]int 
	head int 
	tail int 
}
func (this *CircleQueue)IsFull()bool{
	return (this.tail + 1)%this.maxSize == this.head
}
func(this *CircleQueue)IsEmpty()bool{
	return this.tail == this.head 
}
func (this *CircleQueue)Size()int {       //关键！！！
	return (this.tail + this.maxSize - this.head)%this.maxSize 
}

func (this *CircleQueue)Push(val int)(err error){
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.array[this.tail]=val 
	this.tail=(this.tail+1)%this.maxSize
	return
}
func (this *CircleQueue)Pop()(val int,err error) {
	if this.IsEmpty(){
		return 0,errors.New("queue empty")
	}
	val = this.array[this.head]
	this.head=(this.head + 1) % this.maxSize
	return
}
func (this *CircleQueue) ListQueue(){
	fmt.Println("环形队列情况如下:")
	size := this.Size()
	if size==0 {
		fmt.Println("队列为空")
	}
	tempHead := this.head 
	for i:=0;i<size;i++ {
		fmt.Printf("arr[%d]=%d\t",tempHead,this.array[tempHead])
		tempHead=(tempHead + 1)%this.maxSize
	}
	fmt.Println()
}

func main(){
	queue := &CircleQueue {
		maxSize:5,
		head:0,
		tail:0,
	}
	var key string 
	var val int 
	for{
		fmt.Println("1. 输入 add 表示添加数据到队列")
		fmt.Println("2. 输入 get 表示从队列获取数据")
		fmt.Println("3. 输入 show 表示显示队列")
		fmt.Println("4. 输入 exit 表示退出队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			}else {
				fmt.Println("加入队列OK")
			}
		case "get":
			val,err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			}else{
				fmt.Println("从队列中取出一个数=",val)
			}
		case "show":
			queue.ListQueue()
		case "exit":
			os.Exit(0)
		}
	}
}

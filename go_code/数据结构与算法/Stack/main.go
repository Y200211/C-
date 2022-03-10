package main
import(
	"fmt"
	"errors"
)
type Stack struct {
	MaxTop int 
	Top int 
	arr [5]int 
}
func (this *Stack) Push(val int) (err error) {
	if this.Top == this.MaxTop - 1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	this.arr[this.Top] = val 
	return 
}
func (this *Stack) Pop() (val int, err error){
	if this.Top == -1{
		fmt.Println("stack empty!")
		return 0,errors.New("stack empty")
	}
	val = this.arr[this.Top]
	this.Top--
	return val,nil 
}
func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("stack empty")
		return 
	}
	fmt.Println("栈的情况如下:")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d\n", i , this.arr[i])
	}
}
func main(){
	stack := &Stack{
		MaxTop : 5,
		Top : -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)

	stack.List()
}
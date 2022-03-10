/*
MyCircularQueue(k): 构造器，设置队列长度为 k 。
Front: 从队首获取元素。如果队列为空，返回 -1 。
Rear: 获取队尾元素。如果队列为空，返回 -1 。
enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
isEmpty(): 检查循环队列是否为空。
isFull(): 检查循环队列是否已满。
*/
type MyCircularQueue struct {
	head int
	tail int
	list[] int 
	size int 
	}
	
	
	func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head : 0,
		tail : 0,
		list : make([]int,k+1,k+1),
		size : k+1,
	
	}
	}
	
	
	func (this *MyCircularQueue) EnQueue(value int) bool {
		
		if this.IsFull() {
			return false
		}
		this.list[this.tail] = value 
		this.tail=(this.tail+1+this.size)%this.size
		return true
	
	}
	
	
	func (this *MyCircularQueue) DeQueue() bool {
		if this.IsEmpty() {
			return false
		}
	
		this.head=(this.head+1+this.size)%this.size
	
	
		return true
	}
	
	
	func (this *MyCircularQueue) Front() int {
		if this.IsEmpty() {
			return -1
		}else{
			return this.list[this.head] 
		}
		
	}
	
	
	func (this *MyCircularQueue) Rear() int {
		if this.IsEmpty() {
			return -1
		}else{
		return this.list[(this.tail-1+this.size)%this.size]
		}
	}
	
	
	func (this *MyCircularQueue) IsEmpty() bool {
	return (this.head) == (this.tail) 
	}
	
	func (this *MyCircularQueue) IsFull() bool {
		return (this.tail + 1) % this.size == this.head
	}
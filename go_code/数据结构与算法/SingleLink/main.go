package main
import(
	"fmt"
)
type HeroNode struct{
	no int
	name string 
	nickname string 
	next *HeroNode//标识指向下一个节点
}
//方法一，直接来吧【简单无排序】
func InsertHeroNode(head *HeroNode,newHeroNode *HeroNode) {
	temp := head 
	for {
		if temp.next == nil{
			break
		}
		temp = temp.next 
	}
	temp.next = newHeroNode

}
//方法二，根据no编号从大到小插入【实用，能按序号排好】
//显示链表内所有信息
func InsertHeroNode2(head *HeroNode,newHeroNode *HeroNode){
	temp := head 
	flag := true
	for {
		if temp.next == nil {
			break
		} else if temp.next.no>=newHeroNode.no {
			break 
		}else if temp.next.no == newHeroNode.no{
			flag = false
			break 
		}
		temp = temp.next 
	}
	if !flag {
		fmt.Println("对不起，已经存在 no=",newHeroNode.no)
		return 
	} else {
		newHeroNode.next = temp.next 
		temp.next = newHeroNode
	}
}

func ListHeroNode (head *HeroNode){
	temp := head 
	if temp.next == nil {
		fmt.Println("空空如也。。。")
		return
	}
	for {
		fmt.Printf("[%d,%s,%s]==>",temp.next.no,temp.next.name,temp.next.nickname)
	
	temp = temp.next 
	if temp.next == nil {
		break
	}
}
}
func main(){
	head := &HeroNode{}
	hero1 := &HeroNode {
		no : 1,
		name : "宋江",
		nickname:"及时雨",
	}
	
	hero2 := &HeroNode {
		no : 2,
		name : "卢俊义",
		nickname:"玉麒麟",
	}
	hero3 := &HeroNode {
		no : 3,
		name : "林冲",
		nickname:"豹子头",
	}
	hero4 := &HeroNode {
		no : 4,
		name : "吴用",
		nickname:"智多星",
	}

	InsertHeroNode2(head,hero3)
	InsertHeroNode2(head,hero1)
	InsertHeroNode2(head,hero2)
	InsertHeroNode2(head,hero4)

	ListHeroNode(head)

}
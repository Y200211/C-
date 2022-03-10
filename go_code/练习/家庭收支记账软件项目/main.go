package main
import(
	"fmt"
	
)
func main(){
	var key,sum,add,min int
	sum = 10000
	var cause string
	loop := true
	details := ""
	for{
	fmt.Println("----------------家庭收支记账软件----------------")
	fmt.Println("")
	fmt.Println("                1 收支明细")
	fmt.Println("                2 登记收入")
	fmt.Println("                3 登记支出")
	fmt.Println("                4 退出")
	fmt.Println("")
	fmt.Printf("                请选择(1-4):")
	fmt.Scan(&key)
	switch key {
	case 1 :
		fmt.Println("----------------家庭收支明细记录----------------")
		fmt.Printf("收支\t账户金额\t收支金额\t说  明\n")
		fmt.Println(details)
		
		
	case 2 :
		fmt.Println("本次收入金额：")
		fmt.Scan(&add)
		sum += add
		fmt.Println("本次收入说明：")
		fmt.Scan(&cause)
		details += fmt.Sprintf("\n收入\t%v\t%v\t%v",sum,add,cause)
	case 3 :
		fmt.Println("本次支出金额：")
		fmt.Scan(&min)
		sum -= min
		fmt.Println("本次支出说明：")
		fmt.Scan(&cause)
		details += fmt.Sprintf("\n支出\t%v\t    %v    \t%v",sum,min,cause)
	case 4 :
		loop = false
		fmt.Println("             已退出家庭收支记账软件")
	default:
		fmt.Println("请输入正确的选项...")
	}
	if !loop {
		break
	}
	fmt.Println("")
	fmt.Println("----------------当前收支明细记录----------------")
	fmt.Println("收支\t账户金额\t收支金额\t说  明")
	fmt.Println(details)

	}

}
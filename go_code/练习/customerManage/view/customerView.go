package main
import (
	"fmt"
	"go_code/customerManage/service"
)
type customerView struct{
	//定义必要的字段
	key string 
	loop bool 
	customerService *service.CustomerService
}

func (this *customerView) list() {
	customers := this.customerService.List()
	fmt.Println("-----------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0 ; i<len(customers); i++{
	
	fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
		fmt.Println("------------------客户列表完成--------------------")
}


func (this *customerView) mainMenu() {
	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1—5): ")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1" :
			fmt.Println("添 加 客 户")
		case "2" :
			fmt.Println("修 改 客 户")
		case "3" :
			fmt.Println("删 除 客 户")
		case "4" :
			this.list()
		case "5" :
			this.loop = false
			default :
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}

func  main(){
	//在主函数中，创建一个customerView,并运行显示主菜单
	customerView := customerView{
		key : "",
		loop : true,
	}
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
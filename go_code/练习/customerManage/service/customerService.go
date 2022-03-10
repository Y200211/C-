package service
import (
	"go_code/customerManage/model"

)

type CustomerService struct {
	customers []model.Customer
	customerNum int 
}

//编写一个方法，可以返回*CustomerService
func NewCustomerService() *CustomerService {
	//为了能够看到有客户在切片中，初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer (1,"张三","男",20,"112","zs@sohu.com")
	customerService.customers = append(customerService.customers,customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer{
	return this.customers 
}
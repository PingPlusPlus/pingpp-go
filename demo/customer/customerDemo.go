package customer

import (
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/customer"
)

var Demo = new(CustomerDemo)

type CustomerDemo struct {
	demoAppID string
}

func (c *CustomerDemo) Setup(app string) {
	c.demoAppID = app
}

func (c *CustomerDemo) New() (*pingpp.Customer, error) {
	sms_code := make(map[string]interface{})
	sms_code["id"] = "sms_BDIQG8JZnQhUN1hAbwYubhP3"
	sms_code["code"] = "123222"

	param := &pingpp.CustomerParams{
		App:         c.demoAppID,
		Source:      "tok_ALeWEHQEp1wk9Ebep6a2EhVy",
		Sms_code:    sms_code,
		Description: "create test customer",
		Email:       "newcustomer@test.com",
	}

	return customer.New(param)
}

func (c *CustomerDemo) Get() (*pingpp.Customer, error) {
	return customer.Get("cus_ALeWGZ8lsN9Czk")
}

func (c *CustomerDemo) List() *customer.Iter {
	params := &pingpp.CustomerListParams{}
	params.Filters.AddFilter("limit", "", "1")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	return customer.List(params)
}

func (c *CustomerDemo) Update() (*pingpp.Customer, error) {
	cus_id := "cus_ALeWGZ8lsN9Czk"

	metadata := make(map[string]interface{})
	metadata["red"] = "yello"

	param := &pingpp.CustomerUpdateParams{
		// Default_source: "card_yTWjr1eTWznH8Ci1CC00SWf5",
		Description: "update test customer",
		Email:       "updatecustomer@test.com",
		Metadata:    metadata,
	}

	return customer.Update(cus_id, param)
}

func (c *CustomerDemo) Delete() (map[string]interface{}, error) {
	return customer.Delete("cus_ALeWGZ8lsN9Czk")
}

func (c *CustomerDemo) Run() {
	c.New()
	c.Get()
	c.List()
	c.Update()
	c.Delete()
}

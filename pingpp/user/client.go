package user

import (
	"fmt"
	"log"
	"net/url"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

func New(appId string, params *pingpp.UserParams) (*pingpp.User, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *pingpp.UserParams) (*pingpp.User, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("UserParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	user := &pingpp.User{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/users", appId), c.Key, nil, paramsString, user)
	return user, err
}

func Get(appId, userId string) (*pingpp.User, error) {
	return getC().Get(appId, userId)
}

func (c Client) Get(appId, userId string) (*pingpp.User, error) {
	user := &pingpp.User{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users/%s", appId, userId), c.Key, nil, nil, user)
	return user, err
}

func List(appId string, params *pingpp.PagingParams) (*pingpp.UserList, error) {
	return getC().List(appId, params)
}
func (c Client) List(appId string, params *pingpp.PagingParams) (*pingpp.UserList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	userList := &pingpp.UserList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users", appId), c.Key, body, nil, userList)
	return userList, err
}

func Update(appId, userId string, params map[string]interface{}) (*pingpp.User, error) {
	return getC().Update(appId, userId, params)
}

func (c Client) Update(appId, userId string, params map[string]interface{}) (*pingpp.User, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of update user  to pingpp is :\n %v\n ", string(paramsString))
	}

	user := &pingpp.User{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/users/%s", appId, userId), c.Key, nil, paramsString, user)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return user, err
}

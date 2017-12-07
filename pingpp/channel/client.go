package channel

import (
	"fmt"
	"log"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

func New(appId, subAppId string, params *pingpp.ChannelParams) (*pingpp.Channel, error) {
	return getC().New(appId, subAppId, params)
}

func (c Client) New(appId, subAppId string, params *pingpp.ChannelParams) (*pingpp.Channel, error) {
	paramsString, errs := pingpp.JsonEncode(params)
	if errs != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("ChannelParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	channel := &pingpp.Channel{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/sub_apps/%s/channels", appId, subAppId), c.Key, nil, paramsString, channel)
	return channel, err
}

func Get(appId, subAppId, channel string) (*pingpp.Channel, error) {
	return getC().Get(appId, subAppId, channel)
}

func (c Client) Get(appId, subAppId, channelName string) (*pingpp.Channel, error) {
	channel := &pingpp.Channel{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), c.Key, nil, nil, channel)
	return channel, err
}

func Delete(appId, subAppId, channelName string) (*pingpp.ChannelDeleteResult, error) {
	return getC().Delete(appId, subAppId, channelName)
}

func (c Client) Delete(appId, subAppId, channelName string) (*pingpp.ChannelDeleteResult, error) {
	result := &pingpp.ChannelDeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), c.Key, nil, nil, result)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return result, err
}

func Update(appId, subAppId, channelName string, params pingpp.ChannelUpdateParams) (*pingpp.Channel, error) {
	return getC().Update(appId, subAppId, channelName, params)
}

func (c Client) Update(appId, subAppId, channelName string, params pingpp.ChannelUpdateParams) (*pingpp.Channel, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of update Channel  to pingpp is :\n %v\n ", string(paramsString))
	}

	channel := &pingpp.Channel{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), c.Key, nil, paramsString, channel)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return channel, err
}

package cardInfo

import (
	"log"
	"time"

	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
)

// Client cardInfo 请求客户端
type Client struct {
	B   pingpp.Backend
	Key string
}

func getC() Client {
	return Client{pingpp.GetBackend(pingpp.APIBackend), pingpp.Key}
}

// New 发送 /card_info 请求
func New(params *pingpp.CardInfoParams) (*pingpp.CardInfo, error) {
	return getC().New(params)
}

// New 发送 /card_info 请求
func (c Client) New(params *pingpp.CardInfoParams) (*pingpp.CardInfo, error) {
	start := time.Now()
	paramsString, err := pingpp.JsonEncode(params)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("CardInfoParams Marshall Errors is : %q\n", err)
		}
		return nil, err
	}
	if pingpp.LogLevel > 2 {
		log.Printf("params of cardInfo request to pingpp is :\n %v\n ", string(paramsString))
	}

	cardInfo := &pingpp.CardInfo{}
	errch := c.B.Call("POST", "/card_info", c.Key, nil, paramsString, cardInfo)
	if errch != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	}
	if pingpp.LogLevel > 2 {
		log.Println("CardInfo completed in ", time.Since(start))
	}
	return cardInfo, nil
}

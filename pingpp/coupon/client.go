package coupon

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

//创建优惠券
func New(appId, userId string, params *pingpp.CouponParams) (*pingpp.Coupon, error) {
	return getC().New(appId, userId, params)
}

func (c Client) New(appId, userId string, params *pingpp.CouponParams) (*pingpp.Coupon, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of create coupon request to pingpp is :\n %v\n ", string(paramsString))
	}

	coupon := &pingpp.Coupon{}

	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/users/%s/coupons", appId, userId), c.Key, nil, paramsString, coupon)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return coupon, err
}

//批量创建优惠券
func BatchNew(appId, couponTmplId string, params *pingpp.BatchCouponParams) (*pingpp.CouponList, error) {
	return getC().BatchNew(appId, couponTmplId, params)
}

func (c Client) BatchNew(appId, couponTmplId string, params *pingpp.BatchCouponParams) (*pingpp.CouponList, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of create coupons request to pingpp is :\n %v\n ", string(paramsString))
	}

	couponList := &pingpp.CouponList{}

	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/coupon_templates/%s/coupons", appId, couponTmplId), c.Key, nil, paramsString, couponList)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return couponList, err
}

//更新优惠券
func Update(appId, userId, couponId string, params *pingpp.CouponUpdateParams) (*pingpp.Coupon, error) {
	return getC().Update(appId, userId, couponId, params)
}

func (c Client) Update(appId, userId, couponId string, params *pingpp.CouponUpdateParams) (*pingpp.Coupon, error) {
	paramsString, _ := pingpp.JsonEncode(params)
	if pingpp.LogLevel > 2 {
		log.Printf("params of update coupon  to pingpp is :\n %v\n ", string(paramsString))
	}

	coupon := &pingpp.Coupon{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/users/%s/coupons/%s", appId, userId, couponId), c.Key, nil, paramsString, coupon)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return coupon, err
}

//删除优惠券
func Delete(appId, userId, couponId string) (*pingpp.DeleteResult, error) {
	return getC().Delete(appId, userId, couponId)
}

func (c Client) Delete(appId, userId, couponId string) (*pingpp.DeleteResult, error) {
	result := &pingpp.DeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/apps/%s/users/%s/coupons/%s", appId, userId, couponId), c.Key, nil, nil, result)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Delete Coupon Template error: %v\n", err)
		}
	}
	return result, err
}

//查询指定的优惠券模板
func Get(appId, userId, couponId string) (*pingpp.Coupon, error) {
	return getC().Get(appId, userId, couponId)
}

func (c Client) Get(appId, userId, couponId string) (*pingpp.Coupon, error) {
	var body *url.Values
	body = &url.Values{}
	coupon := &pingpp.Coupon{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users/%s/coupons/%s", appId, userId, couponId), c.Key, body, nil, coupon)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Coupon error: %v\n", err)
		}
	}
	return coupon, err
}

//用户的优惠券列表
func UserList(appId, userId string, params *pingpp.PagingParams) (*pingpp.CouponList, error) {
	return getC().UserList(appId, userId, params)
}

func (c Client) UserList(appId, userId string, params *pingpp.PagingParams) (*pingpp.CouponList, error) {
	body := &url.Values{}
	params.Filters.AppendTo(body)

	couponList := &pingpp.CouponList{}
	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/users/%s/coupons", appId, userId), c.Key, body, nil, couponList)
	if err != nil {
		if pingpp.LogLevel > 0 {
			log.Printf("Get Coupon error: %v\n", err)
		}
	}
	return couponList, err
}

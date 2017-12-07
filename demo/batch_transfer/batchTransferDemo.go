/* *
 * Ping++ Server SDK
 * 说明：
 * 以下代码只是为了方便商户测试而提供的样例代码，商户可以根据自己网站的需要，按照技术文档编写, 并非一定要使用该代码。
 * 该代码仅供学习和研究 Ping++ SDK 使用，只是提供一个参考。
 */
package batch_transfer

import (
	"time"

	"github.com/pingplusplus/pingpp-go/demo/common"
	pingpp "github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/batchTransfer"
)

var Demo = new(BatchTransfer)

type BatchTransfer struct {
	demoAppID           string
	demoBatchTransferID string
}

func (c *BatchTransfer) Setup(app string) {
	c.demoAppID = app
}

//创建 Batch transfer 对象-unionpay渠道
func (c *BatchTransfer) Unionpay_new() (*pingpp.BatchTransfer, error) {
	recipients := []pingpp.BatchTransferRecipient{
		{Account: "6214850266666666", Amount: 5000, Name: "张三", Description: "Your description", Open_bank: "招商银行", Open_bank_code: "0308"},
		{Account: "6214850288888888", Amount: 3000, Name: "李四", Description: "Your description", Open_bank: "招商银行", Open_bank_code: "0308"},
	}
	params := &pingpp.BatchTransferParams{
		App:         c.demoAppID,
		Amount:      8000,
		Batch_no:    time.Now().Format("060102150405"),
		Channel:     "unionpay",
		Description: "Your description",
		Type:        "b2c",
		Recipients:  recipients,
	}

	return batchTransfer.New(params)
}

// 创建 Batch transfer 对象-alipay渠道
func (c *BatchTransfer) Alipay_new() (*pingpp.BatchTransfer, error) {
	recipients := []pingpp.BatchTransferRecipient{
		{Account: "account01@alipay.com", Amount: 5000, Name: "张三", Description: "Your description"},
		{Account: "account02@alipay.com", Amount: 3000, Name: "李四", Description: "Your description"},
	}
	params := &pingpp.BatchTransferParams{
		App:         c.demoAppID,
		Amount:      8000,
		Batch_no:    time.Now().Format("060102150405"),
		Channel:     "alipay",
		Description: "Your description",
		Type:        "b2c",
		Recipients:  recipients,
	}

	return batchTransfer.New(params)
}

//创建 Batch transfer 对象-wx_pub渠道
func (c *BatchTransfer) Wxpub_new() (*pingpp.BatchTransfer, error) {
	recipients := []pingpp.BatchTransferRecipient{
		{Account: "656565656565656565656565", Amount: 5000, Name: "张三", Description: "Your description", ForceCheck: false},
		{Account: "585858585858585858585858", Amount: 3000, Name: "张三", Description: "Your description", ForceCheck: false},
	}
	params := &pingpp.BatchTransferParams{
		App:         c.demoAppID,
		Amount:      8000,
		Batch_no:    time.Now().Format("060102150405"),
		Channel:     "wx_pub",
		Description: "Your description",
		Type:        "b2c",
		Recipients:  recipients,
	}

	return batchTransfer.New(params)
}

//创建 Batch transfer 对象-allinpay渠道
func (c *BatchTransfer) Allinpay_new() (*pingpp.BatchTransfer, error) {
	recipients := []pingpp.BatchTransferRecipient{
		{Account: "656565656565656565656565", Amount: 5000, Name: "张三", Description: "Your description", Open_bank_code: "0308", BusinussCode: "12223", CardType: 1},
		{Account: "585858585858585858585858", Amount: 3000, Name: "李四", Description: "Your description", Open_bank_code: "0308", BusinussCode: "12223"},
	}
	params := &pingpp.BatchTransferParams{
		App:         c.demoAppID,
		Amount:      8000,
		Batch_no:    time.Now().Format("060102150405"),
		Channel:     "allinpay",
		Description: "Your description",
		Type:        "b2c",
		Recipients:  recipients,
	}

	return batchTransfer.New(params)
}

//查询 Batch transfer 对象
func (c *BatchTransfer) Get() (*pingpp.BatchTransfer, error) {
	return batchTransfer.Get(c.demoBatchTransferID)
}

//查询 Batch transfer 对象列表
//更多查询参数可以参照此链接 https://www.pingxx.com/api?language=Go#查询-batch-transfer-对象列表
func (c *BatchTransfer) List() (*pingpp.BatchTransferlList, error) {
	params := &pingpp.PagingParams{}
	params.Filters.AddFilter("page", "", "1")     //页码，取值范围：1~1000000000；默认值为"1"
	params.Filters.AddFilter("per_page", "", "2") //每页数量，取值范围：1～100；默认值为"20"
	params.Filters.AddFilter("app", "", c.demoAppID)
	return batchTransfer.List(params)
}

// 取消付款 (仅unionpay渠道支持)
// unionpay 渠道在 batch transfer 对象请求成功后，延时5分钟发送转账，5分钟内订单处于scheduled的准备发送状态，且可调用该接口通过 batch transfer 对象的 id 更新一个已创建的 batch transfer 对象，即取消该笔转账
func (c *BatchTransfer) Cancel() (*pingpp.BatchTransfer, error) {
	return batchTransfer.Cancel(c.demoBatchTransferID) // 批量转账对象id ，由 Ping++ 生成（必须是unionpay渠道）
}

func (c *BatchTransfer) Run() {
	batch_transfer_unionpay, err := c.Unionpay_new()
	common.Response(batch_transfer_unionpay, err)
	c.demoBatchTransferID = batch_transfer_unionpay.Id
	batch_transfer, err := c.Get()
	common.Response(batch_transfer, err)
	batch_transfer_list, err := c.List()
	common.Response(batch_transfer_list, err)
}

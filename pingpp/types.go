// ping++ golang sdk 数据类型定义
// types涵盖了以下数据格式
//    1.支付订单对象Charge;
//    2.订单退款对象Refund;
//    3.红包对象RedEnvelope;
//    4.企业转账对象Transfer;
//    5.应用内快捷支付对象Card/Customer/Token;
package pingpp

//应用信息数据类型
type App struct {
	Id string `json:"id"`
}

type Data struct {
	Object map[string]interface{} `json:"object"`
}

// 数据列表元数据类型
type ListMeta struct {
	Object string `json:"object"`
	More   bool   `json:"has_more"`
	URL    string `json:"url"`
}

/*支付相关数据类型*/
type (
	// 支付请求数据类型
	ChargeParams struct {
		Order_no    string                 `json:"order_no"`
		App         App                    `json:"app"`
		Channel     string                 `json:"channel"`
		Amount      uint64                 `json:"amount"`
		Currency    string                 `json:"currency"`
		Client_ip   string                 `json:"client_ip"`
		Subject     string                 `json:"subject"`
		Body        string                 `json:"body"`
		Extra       map[string]interface{} `json:"extra,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Time_expire uint64                 `json:"time_expire,omitempty"`
		Description string                 `json:"description,omitempty"`
	}
	// Charge列表查询请求 数据类型
	ChargeListParams struct {
		ListParams
		Created int64
	}

	// Charge数据类型
	Charge struct {
		ID              string                 `json:"id"`
		Object          string                 `json:"object"`
		Created         int64                  `json:"created"`
		Livemode        bool                   `json:"livemode"`
		Paid            bool                   `json:"paid"`
		Refunded        bool                   `json:"refunded"`
		App             string                 `json:"app"`
		Channel         string                 `json:"channel"`
		Order_no        string                 `json:"order_no"`
		Client_ip       string                 `json:"client_ip"`
		Amount          uint64                 `json:"amount"`
		Amount_settle   uint64                 `json:"amount_settle"`
		Currency        string                 `json:"currency"`
		Subject         string                 `json:"subject"`
		Body            string                 `json:"body"`
		Extra           map[string]interface{} `json:"extra"`
		Time_paid       uint64                 `json:"time_paid"`
		Time_expire     uint64                 `json:"time_expire"`
		Time_settle     uint64                 `json:"time_settle"`
		Transaction_no  string                 `json:"transaction_no"`
		Refunds         *RefundList            `json:"refunds"`
		Amount_refunded uint64                 `json:"amount_refunded"`
		Failure_code    string                 `json:"failure_code"`
		Failure_msg     string                 `json:"failure_msg"`
		Metadata        map[string]interface{} `json:"metadata"`
		Credential      map[string]interface{} `json:"credential"`
		Description     string                 `json:"description"`
	}

	// Charge列表数据类型
	ChargeList struct {
		ListMeta
		Values []*Charge `json:"data"`
	}
)

/*退款数据类型*/
type (
	// 退款请求数据类型
	RefundParams struct {
		Amount      uint64                 `json:"amount,omitempty"`
		Description string                 `json:"description"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
	}

	// 退款查询请求的数据类型
	RefundListParams struct {
		ListParams
		Charge string
	}

	// 付款退款数据类型
	Refund struct {
		ID           string                 `json:"id"`
		Object       string                 `json:"object"`
		Order_no     string                 `json:"order_no"`
		Amount       uint64                 `json:"amount"`
		Succeed      bool                   `json:"succeed"`
		Status       string                 `json:"status"`
		Created      uint64                 `json:"created"`
		Time_succeed uint64                 `json:"time_succeed"`
		Description  string                 `json:"description"`
		Failure_code string                 `json:"failure_code"`
		Failure_msg  string                 `json:"failure_msg"`
		Metadata     map[string]interface{} `json:"metadata"`
		Charge_id    string                 `json:"charge"`
	}
	// 付款查询结果列表数据类型
	RefundList struct {
		ListMeta
		Values []*Refund `json:"data"`
	}
)

/*红包请求数据类型*/
type (
	// 红包请求的数据类型
	RedEnvelopeParams struct {
		App         App                    `json:"app"`
		Channel     string                 `json:"channel"`
		Order_no    string                 `json:"order_no"`
		Amount      uint64                 `json:"amount"`
		Currency    string                 `json:"currency"`
		Recipient   string                 `json:"recipient"`
		Subject     string                 `json:"subject"`
		Body        string                 `json:"body"`
		Description string                 `json:"description"`
		Extra       map[string]interface{} `json:"extra"`
	}

	RedEnvelopeListParams struct {
		ListParams
		Created int64
	}

	// 红包数据类型
	RedEnvelope struct {
		Id             string                 `json:"id"`
		Object         string                 `json:"object"`
		Created        uint64                 `json:"created"`
		Received       uint64                 `json:"received"`
		Livemode       bool                   `json:"livemode"`
		Status         string                 `json:"status"`
		App            string                 `json:"app"`
		Channel        string                 `json:"channel"`
		Order_no       string                 `json:"order_no"`
		Transaction_no string                 `json:"transaction_no"`
		Amount         uint64                 `json:"amount"`
		Currency       string                 `json:"currency"`
		Recipient      string                 `json:"recipient"`
		Subject        string                 `json:"subject"`
		Body           string                 `json:"body"`
		Description    string                 `json:"description"`
		Failure_msg    string                 `json:"failure_msg"`
		Extra          map[string]interface{} `json:"extra"`
		Metadata       map[string]interface{} `json:"metadata"`
	}

	// 红包查询结果列表数据类型
	RedEnvelopeList struct {
		ListMeta
		Values []*RedEnvelope `json:"data"`
	}
)

/*企业转账*/
type (
	// 企业转账请求数据类型
	TransferParams struct {
		App         App                    `json:"app"`
		Channel     string                 `json:"channel"`
		Order_no    string                 `json:"order_no"`
		Amount      uint64                 `json:"amount"`
		Type        string                 `json:"type"`
		Currency    string                 `json:"currency"`
		Recipient   string                 `json:"recipient"`
		Description string                 `json:"description"`
		Extra       map[string]interface{} `json:"extra"`
	}

	//企业转账列表查询数据类型
	TransferListParams struct {
		ListParams
		Created int64
	}

	// 企业转账数据类型
	Transfer struct {
		Id              string                 `json:"id"`
		Object          string                 `json:"object"`
		Type            string                 `json:"type"`
		Created         int64                  `json:"created"`
		Time_transfered int64                  `json:"time_transfered"`
		Livemode        bool                   `json:"livemode"`
		Status          string                 `json:"status"`
		App             string                 `json:"app"`
		Channel         string                 `json:"channel"`
		Order_no        string                 `json:"order_no"`
		Amount          uint64                 `json:"amount"`
		Currency        string                 `json:"currency"`
		Recipient       string                 `json:"recipient"`
		Description     string                 `json:"description"`
		Transaction_no  string                 `json:"transaction_no"`
		Failure_msg     string                 `json:"failure_msg"`
		Extra           map[string]interface{} `json:"extra"`
		Metadata        map[string]interface{} `json:"metadata"`
	}
	// 企业转账列表数据类型
	TransferList struct {
		ListMeta
		Values []*Transfer `json:"data"`
	}
)

/*应用内快捷支付相关数据类型*/
type (

	//创建顾客的请求参数
	CustomerParams struct {
		App         string                 `json:"app"`
		Source      interface{}            `json:"source"`
		Sms_id      string                 `json:"sms_id"`
		Sms_code    string                 `json:"sms_code"`
		Description string                 `json:"description,omitempty"`
		Email       string                 `json:"email,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
	}

	//更新顾客信息的请求参数
	CustomerUpdateParams struct {
		Description    string                 `json:"description,omitempty"`
		Email          string                 `json:"email,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		Default_source string                 `json:"default_source,omitempty"`
	}

	//查询顾客的请求参数
	CustomerListParams struct {
		ListParams
		Created int64
	}

	//顾客列表数据类型
	CustomerList struct {
		ListMeta
		Values []*Customer `json:"data"`
	}

	//顾客信息数据类型
	Customer struct {
		ID             string                 `json:"id"`
		Object         string                 `json:"object"`
		Created        int64                  `json:"created"`
		Livemode       bool                   `json:"livemode"`
		App            string                 `json:"app"`
		Name           string                 `json:"name"`
		Email          string                 `json:"email"`
		Currency       string                 `json:"currency"`
		Description    string                 `json:"description"`
		Metadata       map[string]interface{} `json:"metadata"`
		Source         *CardList              `json:"sources"`
		Default_source string                 `json:"default_source"`
	}

	//创建 Card 对象的请求参数
	CardParams struct {
		Source   interface{} `json:"source"`
		Sms_id   string      `json:"sms_id"`
		Sms_code string      `json:"sms_code"`
	}

	//查询 Card 对象的请求参数
	CardListParams struct {
		ListParams
		Created int64
	}

	//Card 对象列表数据类型
	CardList struct {
		ListMeta
		Values []*Card `json:"data"`
	}

	//Card 对象数据类型
	Card struct {
		ID       string `json:"id"`
		Object   string `json:"object"`
		Created  int64  `json:"created"`
		Last4    string `json:"last4"`
		Funding  string `json:"funding"`
		Brand    string `json:"brand"`
		Bank     string `json:"bank"`
		Customer string `json:"customer"`
	}

	//查询 Token 对象的请求参数
	TokenParams struct {
		Order_no   string      `json:"order_no"`
		Amount     uint64      `json:"amount"`
		App        string      `json:"app"`
		//Attachable bool        `json:"attachable"`
		Card       interface{} `json:"card"`
	}

	//Token 对象包含card信息
	Token struct {
		ID         string                 `json:"id"`
		Object     string                 `json:"object"`
		Created    int64                  `json:"created"`
		Livemode   bool                   `json:"livemode"`
		Used       bool                   `json:"used"`
		Time_used  int64                  `json:"time_used"`
		//Attachable bool                   `json:"attachable"`
		Type       string                 `json:"type"`
		Card       map[string]interface{} `json:"card"`
		Sms_code   map[string]interface{} `json:"sms_code"`
	}
)

/*webhooks 相关数据类型*/
type (

	// webhooks 反馈数据类型
	Event struct {
		Id               string `json:"id"`
		Created          int64  `json:"created"`
		Livemode         bool   `json:"livemode"`
		Type             string `json:"type"`
		Data             Data   `json:"data"`
		Object           string `json:"object"`
		Pending_webhooks int    `json:"pending_webhooks"`
		Request          string `json:"request"`
	}

	// webhooks 列表查询数据类型
	EventListParams struct {
		ListParams
		Created int64
	}

	// webhooks 列表数据类型
	EventList struct {
		ListMeta
		Values []*Event `json:"data"`
	}

	//webhooks 汇总数据
	Summary struct {
		Acct_id           string `json:"acct_id,omitempty"`
		App_id            string `json:"app_id.omitempty"`
		Acct_display_name string `json:"acct_display_name"`
		App_display_name  string `json:"app_display_name"`
		Summary_from      uint64 `json:"summary_from"`
		Summary_to        uint64 `json:"summary_to"`
		Charges_amount    uint64 `json:"charges_amount"`
		Charges_count     uint64 `json:"charges_count"`
	}
)

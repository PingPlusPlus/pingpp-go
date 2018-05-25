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

//删除接口返回值
type DeleteResult struct {
	Deleted bool   `json:"deleted"`
	ID      string `json:"id"`
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
		Time_expire int64                  `json:"time_expire,omitempty"`
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
		Reversed        bool                   `json:"reversed"`
		App             string                 `json:"app"`
		Channel         string                 `json:"channel"`
		Order_no        string                 `json:"order_no"`
		Client_ip       string                 `json:"client_ip"`
		Amount          uint64                 `json:"amount"`
		Amount_settle   int64                  `json:"amount_settle"`
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
		Amount         uint64                 `json:"amount,omitempty"`
		Description    string                 `json:"description"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		Funding_source string                 `json:"funding_source,omitempty"`
	}

	// 退款查询请求的数据类型
	RefundListParams struct {
		ListParams
		Charge string
	}

	// 付款退款数据类型
	Refund struct {
		ID              string                 `json:"id"`
		Object          string                 `json:"object"`
		Order_no        string                 `json:"order_no"`
		Amount          uint64                 `json:"amount"`
		Succeed         bool                   `json:"succeed"`
		Status          string                 `json:"status"`
		Created         uint64                 `json:"created"`
		Time_succeed    uint64                 `json:"time_succeed"`
		Description     string                 `json:"description"`
		Failure_code    string                 `json:"failure_code"`
		Failure_msg     string                 `json:"failure_msg"`
		Metadata        map[string]interface{} `json:"metadata"`
		Charge_id       string                 `json:"charge"`
		Charge_order_no string                 `json:"charge_order_no"`
		Transaction_no  string                 `json:"transaction_no"`
		Funding_source  string                 `json:"funding_source,omitempty"`
		Extra           map[string]interface{} `json:"extra"`
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
		Refunded       uint64                 `json:"refunded"`
		Livemode       bool                   `json:"livemode"`
		Status         string                 `json:"status"`
		App            string                 `json:"app"`
		Channel        string                 `json:"channel"`
		Order_no       string                 `json:"order_no"`
		Transaction_no string                 `json:"transaction_no"`
		Amount         uint64                 `json:"amount"`
		Amount_settle  uint64                 `json:"amount_settle"`
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
		Description string                 `json:"description"`
		Recipient   string                 `json:"recipient,omitempty"`
		Extra       map[string]interface{} `json:"extra,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
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
		Batch_no        string                 `json:"batch_no"`
		Amount          uint64                 `json:"amount"`
		Amount_settle   int64                  `json:"amount_settle"`
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
		Sms_code    map[string]interface{} `json:"sms_code"`
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
		Source   interface{}            `json:"source"`
		Sms_code map[string]interface{} `json:"sms_code"`
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

	// CardInfoParams 创建 CardInfo 对象的请求参数
	CardInfoParams struct {
		App         string `json:"app"`          // ping++ app
		BankAccount string `json:"bank_account"` // 银行卡号
	}

	// CardInfo 对象数据类型
	CardInfo struct {
		App             string   `json:"app"`              // ping++ app
		CardBin         string   `json:"card_bin"`         // 卡bin信息
		CardType        int      `json:"card_type"`        // 银行卡号类型，0：借记卡；1：存折；2：信用卡；3：准贷记卡；4：其他。jdpay 不支持 1：存折
		OpenBankCode    string   `json:"open_bank_code"`   // 机构联行号
		OpenBank        string   `json:"open_bank"`        // 所属银行名称
		SupportChannels []string `json:"support_channels"` // 支持的出款渠道名称列表。可能包含的值有 "unionpay"、"unionpay_gz"、"allinpay"、"jdpay"
	}

	//查询 Token 对象的请求参数
	TokenParams struct {
		Order_no string `json:"order_no"`
		Amount   uint64 `json:"amount"`
		App      string `json:"app"`
		//Attachable bool        `json:"attachable"`
		Card interface{} `json:"card"`
	}

	//Token 对象包含card信息
	Token struct {
		ID        string `json:"id"`
		Object    string `json:"object"`
		Created   int64  `json:"created"`
		Livemode  bool   `json:"livemode"`
		Used      bool   `json:"used"`
		Time_used int64  `json:"time_used"`
		//Attachable bool                   `json:"attachable"`
		Type     string                 `json:"type"`
		Card     map[string]interface{} `json:"card"`
		Sms_code map[string]interface{} `json:"sms_code"`
	}

	// AgreementParams 请求签约对象参数
	AgreementParams struct {
		// App ID
		App string `json:"app"`
		// ContractNo 签约协议号
		ContractNo string `json:"contract_no"`
		// Channel 签约渠道
		Channel string `json:"channel"`
		// Extra 附加信息
		Extra map[string]interface{} `json:"extra,omitempty"`
		// Metadata 元数据
		Metadata map[string]interface{} `json:"metadata,omitempty"`
	}
	AgreementUpdateParams struct {
		// Status 状态
		Status string `json:"status,omitempty"`
	}
	// Agreement 签约对象
	Agreement struct {
		// ID 签约对象ID
		ID string `json:"id"`
		// Object 值为 agreement。
		Object string `json:"object"`
		// Livemode 是否是 live 模式。
		Livemode bool `json:"livemode"`
		// App APP_ID
		App string `json:"app"`
		// Created 创建时间，用 Unix 时间戳表示。
		Created int `json:"created"`
		// Channel 签约渠道
		Channel string `json:"channel"`
		// ContractNo [商户]签约协议号。
		ContractNo string `json:"contract_no"`
		// ContractID 渠道签约ID
		ContractID string `json:"contract_id"`
		// Credential 签约渠道凭证
		Credential map[string]interface{} `json:"credential"`
		// Status 签约状态 (created:待签约 ,succeeded:签约成功, canceled:已解约 )
		Status string `json:"status"`
		// TimeSucceeded 签约成功时间，用 Unix 时间戳表示。
		TimeSucceeded int `json:"time_succeeded"`
		// TimeCanceled 解约成功时间,用 Unix 时间戳表示。
		TimeCanceled int `json:"time_canceled"`
		// FailureCode 签约错误码，详见 错误 中的错误码描述。
		FailureCode string `json:"failure_code"`
		// FailureMsg 签约错误消息的描述。
		FailureMsg string `json:"failure_msg"`
		// Extra 附加参数
		Extra map[string]interface{} `json:"extra"`
		// Metadata metadata 元数据
		Metadata map[string]interface{} `json:"metadata"`
	}

	AgreementList struct {
		ListMeta
		Values []*Agreement `json:"data"`
	}

	// OpenidParams 对象包含了查询小程序 openid 信息
	OpenidParams struct {
		AppID     string `json:"app_id"`     // 微信小程序应用唯一标识
		AppSecret string `json:"app_secret"` // 微信小程序应用密钥（注意保密）
		Code      string `json:"code"`       // 授权code, 登录时获取的 code

	}
	// Openid 返回了小程序用户信息
	Openid struct {
		ID         string `json:"openid,omitempty"`      // 用户唯一标识
		SessionKey string `json:"session_key,omitempty"` // 会话密钥
		Unionid    string `json:"unionid,omitempty"`     // 同主体的公众号的唯一用户 id

		ErrCode int    `json:"errcode,omitempty"`
		ErrMsg  string `json:"errmsg,omitempty"`
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

	//webhooks 汇总数据
	Summary struct {
		Acct_id           string `json:"acct_id,omitempty"`
		App_id            string `json:"app_id,omitempty"`
		Acct_display_name string `json:"acct_display_name"`
		App_display_name  string `json:"app_display_name"`
		Summary_from      uint64 `json:"summary_from"`
		Summary_to        uint64 `json:"summary_to"`
		Charges_amount    uint64 `json:"charges_amount"`
		Charges_count     uint64 `json:"charges_count"`
	}
)

//身份认证相关数据结构
type (
	IdentificationParams struct {
		Type string                 `json:"type"`
		App  string                 `json:"app"`
		Data map[string]interface{} `json:"data"`
	}
	IdentificationResult struct {
		Type       string                 `json:"type"`
		App        string                 `json:"app"`
		ResultCode int                    `json:"result_code"`
		Message    string                 `json:"message"`
		Paid       bool                   `json:"paid"`
		Data       map[string]interface{} `json:"data"`
	}
)

type (
	BatchTransferRecipient struct {
		Account        string `json:"account,omitempty"`
		User           string `json:"user,omitempty"`
		Amount         int64  `json:"amount"`
		Name           string `json:"name"`
		Description    string `json:"description,omitempty"`
		Transfer       string `json:"transfer,omitempty"`
		Status         string `json:"status,omitempty"`
		Open_bank      string `json:"open_bank,omitempty"`
		Open_bank_code string `json:"open_bank_code,omitempty"`
		Account_type   string `json:"account_type,omitempty"`
		Fee            int64  `json:"fee,omitempty"`
		Failure_msg    string `json:"failure_msg,omitempty"`
		Order_no       string `json:"order_no,omitempty"`
		Transaction_no string `json:"transaction_no,omitempty"`
		BusinussCode   string `json:"business_code,omitempty"`
		CardType       int    `json:"card_type,omitempty"`
		ForceCheck     bool   `json:"force_check,omitempty"`
	}
	BatchTransfer struct {
		Id             string                   `json:"Id"`
		Object         string                   `json:"object"`
		App            string                   `json:"app"`
		Amount         int64                    `json:"amount"`
		Batch_no       string                   `json:"batch_no"`
		Channel        string                   `json:"channel"`
		Currency       string                   `json:"currency"`
		Created        int64                    `json:"created"`
		Description    string                   `json:"description"`
		Extra          map[string]interface{}   `json:"extra"`
		Failure_msg    string                   `json:"failure_msg"`
		Fee            int64                    `json:"fee"`
		Livemode       bool                     `json:"livemode"`
		Metadata       map[string]interface{}   `json:"metadata"`
		Recipients     []BatchTransferRecipient `json:"recipients"`
		Status         string                   `json:"status"`
		Time_succeeded int64                    `json:"time_succeeded"`
		Type           string                   `json:"type"`
	}

	BatchTransferParams struct {
		App         string                   `json:"app,omitempty"`
		Batch_no    string                   `json:"batch_no,omitempty"`
		Channel     string                   `json:"channel,omitempty"`
		Amount      int64                    `json:"amount"`
		Description string                   `json:"description,omitempty"`
		Metadata    map[string]interface{}   `json:"metadata,omitempty"`
		Recipients  []BatchTransferRecipient `json:"recipients,omitempty"`
		Currency    string                   `json:"currency,omitempty"`
		Type        string                   `json:"type"`
	}

	BatchTransferlList struct {
		ListMeta
		Values []*BatchTransfer `json:"data"`
	}
)

type (
	BatchRefund struct {
		Id          string                 `json:"id"`
		App         string                 `json:"app"`
		Object      string                 `json:"object"`
		Batch_no    string                 `json:"batch_no"`
		Created     int64                  `json:"created"`
		Description string                 `json:"description"`
		Metadata    map[string]interface{} `json:"metadata"`
		Charges     []struct {
			Charge         string  `json:"charge"`
			Amount         *int64  `json:"amount,omitempty"`
			Description    *string `json:"description,omitempty"`
			Status         string  `json:"status"`
			Funding_source *string `json:"funding_source,omitempty"`
		} `json:"charges"`
		Refunds        RefundList `json:"refunds"`
		Refund_url     string     `json:"refund_url"`
		Status         string     `json:"status"`
		Time_succeeded int64      `json:"time_succeeded"`
		Livemode       bool       `json:"livemode"`
	}

	BatchRefundParams struct {
		App         string                   `json:"app"`
		Batch_no    string                   `json:"batch_no"`
		Charges     []map[string]interface{} `json:"charges"`
		Description string                   `json:"description,omitempty"`
		Metadata    map[string]interface{}   `json:"metadata,omitempty"`
	}

	BatchRefundlList struct {
		ListMeta
		Values []*BatchRefund `json:"data"`
	}
)

type (
	Customs struct {
		Id               string                 `json:"id"`
		App              string                 `json:"app"`
		Channel          string                 `json:"channel"`
		Trade_no         string                 `json:"trade_no"`
		Customs_code     string                 `json:"customs_code"`
		Amount           int64                  `json:"amount"`
		Charge           string                 `json:"charge"`
		Transport_amount int64                  `json:"transport_amount"`
		Is_split         bool                   `json:"is_split"`
		Sub_order_no     string                 `json:"sub_order_no"`
		Extra            map[string]interface{} `json:"extra"`
		Object           string                 `json:"object"`
		Created          int64                  `json:"created"`
		Time_succeeded   int64                  `json:"time_succeeded"`
		Status           string                 `json:"status"`
		Failure_code     string                 `json:"failure_code"`
		Failure_msg      string                 `json:"failure_msg"`
		Transaction_no   string                 `json:"transaction_no"`
	}

	CustomsParams struct {
		App              string                 `json:"app"`
		Channel          string                 `json:"channel"`
		Trade_no         string                 `json:"trade_no"`
		Customs_code     string                 `json:"customs_code"`
		Amount           int64                  `json:"amount"`
		Charge           string                 `json:"charge"`
		Transport_amount int64                  `json:"transport_amount"`
		Is_split         bool                   `json:"is_split"`
		Sub_order_no     string                 `json:"sub_order_no"`
		Extra            map[string]interface{} `json:"extra,omitempty"`
	}
)

//用户相关数据结构
type (
	UserParams struct {
		ID       string                 `json:"id"`
		Address  string                 `json:"address,omitempty"`
		Avatar   string                 `json:"avatar,omitempty"`
		Email    string                 `json:"email,omitempty"`
		Gender   string                 `json:"gender,omitempty"`
		Mobile   string                 `json:"mobile,omitempty"`
		Name     string                 `json:"name,omitempty"`
		Metadata map[string]interface{} `json:"metadata,omitempty"`
	}

	User struct {
		ID                   string                 `json:"id"`
		Object               string                 `json:"object"`
		App                  string                 `json:"app"`
		Address              string                 `json:"address"`
		Type                 string                 `json:"type"`        // v1.3 add
		RelatedApp           string                 `json:"related_app"` // v1.3 add
		Available_coupons    int64                  `json:"available_coupons"`
		Avatar               string                 `json:"avatar"`
		Available_balance    int64                  `json:"available_balance"`
		Withdrawable_balance int64                  `json:"withdrawable_balance"`
		Created              int64                  `json:"created"`
		Disabled             bool                   `json:"disabled"`
		Email                string                 `json:"email"`
		Gender               string                 `json:"gender"`
		Identified           bool                   `json:"identified"`
		Livemode             bool                   `json:"livemode"`
		Mobile               string                 `json:"mobile"`
		Name                 string                 `json:"name"`
		Metadata             map[string]interface{} `json:"metadata"`
		SettleAccounts       []SettleAccount        `json:"settle_accounts"`
	}

	SettleAccountParams struct {
		Channel   string                 `json:"channel"`
		Recipient map[string]interface{} `json:"recipient"`
	}

	SettleAccount struct {
		ID         string    `json:"id"`
		Object     string    `json:"object"`
		Create     int64     `json:"created"`
		Channel    string    `json:"channel"`
		Recipients Recipient `json:"recipient"`
	}

	Recipient struct {
		Account      string `json:"account"`
		Name         string `json:"name"`
		Type         string `json:"type,omitempty"`
		OpenBank     string `json:"open_bank,omitempty"`
		OpenBankCode string `json:"open_bank_code,omitempty"`
		ForceCheck   bool   `json:"force_check,omitempty"`
	}

	//结算账户列表
	SettleAccountList struct {
		ListMeta
		Values []*SettleAccount `json:"data"`
	}

	//用户列表
	UserList struct {
		ListMeta
		Values []*User `json:"data"`
	}
)

//多级商户相关数据结构 V1.3 add
type (
	SubAppParams struct {
		DisplayName string                 `json:"display_name"`
		User        string                 `json:"user"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Description string                 `json:"description,omitempty"`
		ParentApp   string                 `json:"parent_app,omitempty"`
	}

	SubApp struct {
		ID               string                 `json:"id"`
		Object           string                 `json:"object"`
		Created          int64                  `json:"created"`
		DisplayName      string                 `json:"display_name"`
		Account          string                 `json:"account"`
		Description      string                 `json:"description"`
		Metadata         map[string]interface{} `json:"metadata"`
		AvailableMethods []string               `json:"available_methods"`
		User             string                 `json:"user"`
		Level            int                    `json:"level"`
		ParentApp        string                 `json:"parent_app"`
	}

	SubAppUpdateParams struct {
		DisplayName string                 `json:"display_name,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Description string                 `json:"description,omitempty"`
		ParentApp   string                 `json:"parent_app,omitempty"`
	}

	//只商户列表
	SubAppList struct {
		ListMeta
		Values []*SubApp `json:"data"`
	}

	Channel struct {
		Object      string                 `json:"object"`
		Created     int64                  `json:"created"`
		Channel     string                 `json:"channel"`
		Params      map[string]interface{} `json:"params"`
		Banned      bool                   `json:"banned"`
		BannedMsg   string                 `json:"banned_msg"`
		Description string                 `json:"description"`
	}

	ChannelParams struct {
		Channel     string                 `json:"channel"`
		Params      map[string]interface{} `json:"params"`
		Banned      bool                   `json:"banned,omitempty"`
		BannedMsg   string                 `json:"banned_msg,omitempty"`
		Description string                 `json:"description,omitempty"`
	}

	ChannelUpdateParams struct {
		Params      map[string]interface{} `json:"params,omitempty"`
		Banned      bool                   `json:"banned,omitempty"`
		BannedMsg   string                 `json:"banned_msg,omitempty"`
		Description string                 `json:"description,omitempty"`
	}

	ChannelDeleteResult struct {
		Deleted bool   `json:"deleted"`
		Channel string `json:"channel"`
	}
)

type (
	CouponTmpl struct {
		ID                   string                 `json:"id"`
		Object               string                 `json:"object"`
		App                  string                 `json:"app"`
		Amount_available     int64                  `json:"amount_available"`
		Amount_off           *int64                 `json:"amount_off"`
		Created              int64                  `json:"created"`
		Expiration           map[string]interface{} `json:"expiration"`
		Livemode             bool                   `json:"livemode"`
		Max_circulation      *int64                 `json:"max_circulation"`
		Max_user_circulation *int64                 `json:"max_user_circulation"`
		Metadata             map[string]interface{} `json:"metadata"`
		Name                 string                 `json:"name"`
		Percent_off          *int64                 `json:"percent_off"`
		Refundable           bool                   `json:"refundable"`
		Times_circulated     int64                  `json:"times_circulated"`
		Times_redeemed       int64                  `json:"times_redeemed"`
		Type                 int64                  `json:"type"`
	}

	CouponTmplParams struct {
		Name             string                 `json:"name,omitempty"`
		Type             int64                  `json:"type"`
		Amount_off       int64                  `json:"amount_off,omitempty"`
		Percent_off      int64                  `json:"percent_off,omitempty"`
		Amount_available int64                  `json:"amount_available,omitempty"`
		Max_circulation  *int64                 `json:"max_circulation,omitempty"`
		Metadata         map[string]interface{} `json:"metadata,omitempty"`
		Expiration       map[string]interface{} `json:"expiration,omitempty"`
		Refundable       *bool                  `json:"refundable,omitempty"`
	}

	CouponTmplUpdateParams struct {
		Metadata map[string]interface{} `json:"metadata,omitempty"`
	}

	//优惠券模板列表
	CouponTmplList struct {
		ListMeta
		Values []*CouponTmpl `json:"data"`
	}
)

type (
	Coupon struct {
		ID                  string                 `json:"id"`
		Object              string                 `json:"object"`
		App                 string                 `json:"app"`
		Actual_amount       *int64                 `json:"actual_amount"`
		Coupon_template     map[string]interface{} `json:"coupon_template"`
		Created             int64                  `json:"created"`
		Livemode            bool                   `json:"livemode"`
		Metadata            map[string]interface{} `json:"metadata"`
		UserTimesCirculated int                    `json:"user_times_circulated"`
		Order               string                 `json:"order"`
		Redeemed            bool                   `json:"redeemed"`
		Time_end            *int64                 `json:"time_end"`
		Time_start          *int64                 `json:"time_start"`
		User                string                 `json:"user"`
		Valid               bool                   `json:"valid"`
	}

	CouponParams struct {
		Coupon_tmpl_id string                 `json:"coupon_template,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
	}

	BatchCouponParams struct {
		Users    []string               `json:"users,omitempty"`
		Metadata map[string]interface{} `json:"metadata,omitempty"`
	}

	CouponUpdateParams struct {
		Metadata map[string]interface{} `json:"metadata"`
	}

	//优惠券列表
	CouponList struct {
		ListMeta
		Values []*Coupon `json:"data"`
	}
)

//Order相关的数据结构
type (
	// 创建Order请求数据类型
	OrderCreateParams struct {
		App               string                 `json:"app"`
		Uid               string                 `json:"uid,omitempty"`
		Merchant_order_no string                 `json:"merchant_order_no,omitempty"`
		Amount            int64                  `json:"amount"`
		Currency          string                 `json:"currency,omitempty"`
		Client_ip         string                 `json:"client_ip,omitempty"`
		Subject           string                 `json:"subject,omitempty"`
		Body              string                 `json:"body,omitempty"`
		Description       string                 `json:"description,omitempty"`
		Coupon            string                 `json:"coupon,omitempty"`
		Actual_amount     int64                  `json:"actual_amount,omitempty"`
		Time_expire       int64                  `json:"time_expire,omitempty"`
		Metadata          map[string]interface{} `json:"metadata,omitempty"`
		ReceiptApp        string                 `json:"receipt_app,omitempty"`
		ServiceApp        string                 `json:"service_app,omitempty"`
		RoyaltyUsers      []RoyaltyUser          `json:"royalty_users,omitempty"`
		RoyaltyTemplate   string                 `json:"royalty_template,omitempty"`
	}
	//分润用户
	RoyaltyUser struct {
		User   string `json:"user"`
		Amount int    `json:"amount"`
	}
	//订单支付请求参数
	OrderPayParams struct {
		Charge_amount   *int64                 `json:"charge_amount,omitempty"`
		Channel         string                 `json:"channel,omitempty"`
		Extra           map[string]interface{} `json:"extra,omitempty"`
		TimeExpire      int64                  `json:"time_expire,omitempty"` // 时间戳
		Charge_order_no string                 `json:"charge_order_no,omitempty"`
	}

	// Order数据类型
	Order struct {
		ID                string                 `json:"id"`
		Object            string                 `json:"object"`
		Created           int64                  `json:"created"`
		Livemode          bool                   `json:"livemode"`
		Refunded          bool                   `json:"refunded"`
		Status            string                 `json:"status"`
		Paid              bool                   `json:"paid"`
		App               string                 `json:"app"`
		Uid               string                 `json:"uid"`
		Merchant_order_no string                 `json:"merchant_order_no"`
		Amount            int64                  `json:"amount"`
		Coupon_amount     int64                  `json:"coupon_amount"`
		Actual_amount     int64                  `json:"actual_amount"`
		Amount_paid       int64                  `json:"amount_paid"`
		Amount_refunded   int64                  `json:"amount_refunded"`
		Currency          string                 `json:"currency"`
		Subject           string                 `json:"subject"`
		Body              string                 `json:"body"`
		Client_ip         string                 `json:"client_ip"`
		Time_paid         int64                  `json:"time_paid"`
		Time_expire       int64                  `json:"time_expire"`
		Coupon            string                 `json:"coupon"`
		Charge            string                 `json:"charge"`
		Charges           ChargeList             `json:"charges"`
		Charge_essentials map[string]interface{} `json:"charge_essentials"`
		Available_balance int64                  `json:"available_balance"`
		ReceiptApp        string                 `json:"receipt_app"`
		ServiceApp        string                 `json:"service_app"`
		AvailableMethods  []string               `json:"available_methods"`
		Description       string                 `json:"description"`
		Metadata          map[string]interface{} `json:"metadata"`
	}

	// Order列表数据类型
	OrderList struct {
		ListMeta
		Values []*Order `json:"data"`
	}

	// Recharge数据类型
	Recharge struct {
		ID                 string                 `json:"id"`
		Object             string                 `json:"object"`
		App                string                 `json:"app"`
		Created            int64                  `json:"created"`
		Livemode           bool                   `json:"livemode"`
		Amount             int                    `json:"amount"`
		Succeeded          bool                   `json:"succeeded"`
		Time_succeed       uint64                 `json:"time_succeed"`
		Refunded           bool                   `json:"refunded"`
		User               string                 `json:"user"`
		From               string                 `json:"from"`
		UserFee            int                    `json:"user_fee"`
		Charge             Charge                 `json:"charge"`
		BalanceBonus       BalanceBonus           `json:"balance_bonus"`
		BalanceTransaction BalanceTransaction     `json:"balance_transaction"`
		Description        string                 `json:"description"`
		Metadata           map[string]interface{} `json:"metadata"`
	}

	// Recharge列表数据类型
	RechargeList struct {
		ListMeta
		Values []*Recharge `json:"data"`
	}

	// 充值退款请求参数
	RechargeRefundParams struct {
		Description    string                 `json:"description,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		Funding_source string                 `json:"funding_source,omitempty"`
	}
)

type (
	//订单退款请求参数
	OrderRefundParams struct {
		Charge         string                 `json:"charge,omitempty"`
		Charge_amount  *int64                 `json:"charge_amount,omitempty"`
		Description    string                 `json:"description,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		Refund_mode    string                 `json:"refund_mode,omitempty"`
		Funding_source string                 `json:"funding_source,omitempty"`
		RoyaltyUsers   []RoyaltyRefundUser    `json:"royalty_users,omitempty"`
	}
	//分润用户
	RoyaltyRefundUser struct {
		User   string `json:"user"`
		Amount int    `json:"amount_refunded"`
	}
	//订单退款对象
	OrderRefund struct {
		ID                  string                 `json:"id"`
		Object              string                 `json:"object"`
		Order               string                 `json:"order"`
		App                 string                 `json:"app"`
		Uid                 string                 `json:"uid"`
		Merchant_order_no   string                 `json:"merchant_order_no"`
		Coupon              string                 `json:"coupon"`
		Amount              int64                  `json:"amount"`
		Coupon_amount       int64                  `json:"coupon_amount"`
		Balance_amount      int64                  `json:"balance_amount"`
		Charge_amount       int64                  `json:"charge_amount"`
		Balance_transaction string                 `json:"balance_transaction"`
		Charge_refund       string                 `json:"charge_refund"`
		Created             int64                  `json:"created"`
		Status              string                 `json:"status"`
		Time_succeed        int64                  `json:"time_succeed"`
		Description         string                 `json:"description"`
		Metadata            map[string]interface{} `json:"metadata"`
		Extra               map[string]interface{} `json:"extra"`
		Refund_mode         string                 `json:"refund_mode"`
	}
	//订单退款列表
	OrderRefundList struct {
		ListMeta
		Values []*OrderRefund `json:"data"`
	}
)

// 分润相关
type (
	Royalty struct {
		ID                 string                 `json:"id"`
		Object             string                 `json:"object"`
		PayerApp           string                 `json:"payer_app"`
		Amount             int                    `json:"amount"`
		Created            int64                  `json:"created"`
		Livemode           bool                   `json:"livemode"`
		Status             string                 `json:"status"`
		Method             string                 `json:"method"`
		RecipientApp       string                 `json:"recipient_app"`
		RoyaltyTransaction string                 `json:"royalty_transaction"`
		RoyaltySettlement  string                 `json:"royalty_settlement"`
		TimeSettled        int64                  `json:"time_settled"`
		SettleAccount      string                 `json:"settle_account"`
		SourceApp          string                 `json:"source_app"`
		SourceUrl          string                 `json:"source_url"`
		SourceNo           string                 `json:"source_no"`
		SourceUser         string                 `json:"source_user"`
		Description        string                 `json:"description"`
		Metadata           map[string]interface{} `json:"metadata"`
	}

	RoyaltyBatchUpdateParams struct {
		Ids         []string               `json:"ids"`
		Method      string                 `json:"method,omitempty"`
		Description string                 `json:"description,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
	}
	// 分润列表数据类型
	RoyaltyList struct {
		ListMeta
		Values []*Royalty `json:"data"`
	}

	RoyaltyTmplList struct {
		ListMeta
		Values []*RoyaltyTmpl `json:"data"`
	}

	RoyaltySettlement struct {
		ID                  string                 `json:"id"`
		Object              string                 `json:"object"`
		PayerApp            string                 `json:"payer_app"`
		Created             int64                  `json:"created"`
		Livemode            bool                   `json:"livemode"`
		Method              string                 `json:"method"`
		Amount              int                    `json:"amount"`
		AmountSucceeded     int                    `json:"amount_succeeded"`
		AmountFailed        int                    `json:"amount_failed"`
		AmountCanceled      int                    `json:"amount_canceled"`
		Count               int                    `json:"count"`
		CountSucceeded      int                    `json:"count_succeeded"`
		CountFailed         int                    `json:"count_failed"`
		CountCanceled       int                    `json:"count_canceled"`
		TimeFinished        int64                  `json:"time_finished"`
		Fee                 int                    `json:"fee"`
		Metadata            map[string]interface{} `json:"metadata"`
		OperationUrl        string                 `json:"operation_url"`
		Status              string                 `json:"status"`
		RoyaltyTransactions RoyaltyTransactionList `json:"royalty_transactions"`
		RoyaltySettlement   string                 `json:"royalty_settlement;omitempty"`
	}

	RoyaltyTmplParams struct {
		App         string `json:"app"`
		Name        string `json:"name,omitempty"`
		Rule        Rule   `json:"rule"`
		Description string `json:"description,omitempty"`
	}

	Rule struct {
		Royalty_mode    string     `json:"royalty_mode,omitempty"`
		Refund_mode     string     `json:"refund_mode,omitempty"`
		Allocation_mode string     `json:"allocation_mode,omitempty"`
		Data            []RuleData `json:"data,omitempty"`
	}

	RuleData struct {
		Level int `json:"level"`
		Value int `json:"value"`
	}

	RoyaltyTmpl struct {
		ID          string `json:"id"`
		Object      string `json:"object"`
		Livemode    bool   `json:"livemode"`
		App         string `json:"app"`
		Name        string `json:"name"`
		Created     int64  `json:"created"`
		Description string `json:"description"`
		Rule        Rule   `json:"rule"`
	}

	RoyaltyTmplUpdateParams struct {
		Rule        Rule   `json:"rule,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}

	RoyaltyTransaction struct {
		ID                string `json:"id"`
		Object            string `json:"object"`
		Amount            int    `json:"amount"`
		Status            string `json:"status"`
		SettleAccount     string `json:"settle_account"`
		RoyaltySettlement string `json:"royalty_settlement"`
		SourceUser        string `json:"source_user"`
		Created           int64  `json:"created"`
		RecipientApp      string `json:"recipient_app"`
		FailureMsg        string `json:"failure_msg"`
		Transfer          string `json:"transfer"`
	}

	RoyaltyTransactionList struct {
		ListMeta
		Values []*RoyaltyTransaction `json:"data"`
	}

	RoyaltySettlementCreateParams struct {
		PayerApp     string                 `json:"payer_app"`
		Method       string                 `json:"method"`
		RecipientApp string                 `json:"recipient_app,omitempty"`
		Created      Created                `json:"created,omitempty"`
		SourceUser   string                 `json:"source_user,omitempty"`
		SourceNo     string                 `json:"source_no,omitempty"`
		MinAmount    int                    `json:"min_amount,omitempty"`
		Metadata     map[string]interface{} `json:"metadata,omitempty"`
		IsPreview    bool                   `json:"is_preview,omitempty"`
	}
	Created struct {
		GT  int64 `json:"gt,omitempty"`
		GTE int64 `json:"gte,omitempty"`
		LT  int64 `json:"lt,omitempty"`
		LTE int64 `json:"lte,omitempty"`
	}
	RoyaltySettlementUpdateParams struct {
		Status string `json:"status"`
	}
	RoyaltySettlementList struct {
		ListMeta
		Values []*RoyaltySettlement `json:"data"`
	}
)

type (
	RechargeParams struct {
		User         string                 `json:"user"`
		Charge       RechargeCharge         `json:"charge"`
		UserFee      int                    `json:"user_fee,omitempty"`
		BalanceBonus RechargeBonus          `json:"balance_bonus,omitempty"`
		FromUser     string                 `json:"from_user,omitempty"`
		Description  string                 `json:"description,omitempty"`
		Metadata     map[string]interface{} `json:"metadata,omitempty"`
	}

	RechargeCharge struct {
		Amount     int    `json:"amount"`
		Channel    string `json:"channel"`
		OrderNo    string `json:"order_no"`
		Subject    string `json:"subject"`
		Body       string `json:"body"`
		TimeExpire int64  `json:"time_expire,omitempty"`
		ClientIp   string `json:"client_ip,omitempty"`
	}
	RechargeBonus struct {
		Amount int `json:"amount,omitempty"`
	}
)

type (

	//企业清算账户交易明细
	// Fee                 int64  `json:"fee"`  v1.3 delete
	// User_fee            int64  `json:"user_fee"` v1.3 delete
	AssetTransaction struct {
		Id                  string `json:"id"`
		Object              string `json:"object"`
		App                 string `json:"app"`
		Amount              int64  `json:"amount"`
		Balance_transaction string `json:"balance_transaction"`
		Created             int64  `json:"created"`
		Description         string `json:"description"`
		Livemode            bool   `json:"livemode"`
		Source              string `json:"source"`
		SourceUrl           string `json:"source_url"` // v1.3 add
		Status              string `json:"status"`
		Time_revoked        int64  `json:"time_revoked"`
		Time_settled        int64  `json:"time_settled"`
		Type                string `json:"type"`
		Method              string `json:"method"`         // v1.3 add
		OrderNo             string `json:"order_no"`       // v1.3 add
		TransactionNo       string `json:"transaction_no"` // v1.3 add
	}

	AssetTransactionList struct {
		ListMeta
		Values []*AssetTransaction `json:"data"`
	}

	AssetStatistic struct {
		Object                    string `json:"object"`
		App                       string `json:"app"`
		Amount_managed            int64  `json:"amount_managed"`
		Amount_settled            int64  `json:"amount_settled"`
		Amount_recharge           int64  `json:"amount_recharge"`
		Amount_payment            int64  `json:"amount_payment"`
		Amount_withdrawal         int64  `json:"amount_withdrawal"`
		Amount_refunded           int64  `json:"amount_refunded"`
		Amount_earning            int64  `json:"amount_earning"`
		Amount_operational        int64  `json:"amount_operational"`
		Fee_total                 int64  `json:"fee_total"`
		Fee_recharge              int64  `json:"fee_recharge"`
		Fee_withdrawal            int64  `json:"fee_withdrawal"`
		User_fee_total            int64  `json:"user_fee_total"`
		User_fee_withdrawal       int64  `json:"user_fee_withdrawal"`
		User_fee_recharge         int64  `json:"user_fee_recharge"`
		User_fee_balance_transfer int64  `json:"user_fee_balance_transfer"`
		Time_start                int64  `json:"time_start"`
		Time_end                  int64  `json:"time_end"`
	}
)

type (
	BalanceTransaction struct {
		Id                string `json:"id"`
		Object            string `json:"object"`
		App               string `json:"app"`
		Amount            int64  `json:"amount"`
		Available_balance int64  `json:"available_balance"`
		Created           int64  `json:"created"`
		Description       string `json:"description"`
		Livemode          bool   `json:"livemode"`
		Source            string `json:"source"`
		Type              string `json:"type"`
		User              string `json:"user"`
	}

	BalanceTransfer struct {
		Id                          string                 `json:"id"`
		Object                      string                 `json:"object"`
		App                         string                 `json:"app"`
		Created                     int64                  `json:"created"`
		Livemode                    bool                   `json:"livemode"`
		Status                      string                 `json:"status"`
		Amount                      int                    `json:"amount"`
		Order_no                    string                 `json:"order_no"`
		User                        string                 `json:"user"`
		Recipient                   string                 `json:"recipient"`
		UserFee                     int                    `json:"user_fee"`
		UserBalanceTransaction      string                 `json:"user_balance_transaction"`
		RecipientBalanceTransaction string                 `json:"recipient_balance_transaction"`
		Description                 string                 `json:"description"`
		Metadata                    map[string]interface{} `json:"metadata,omitempty"`
	}

	BalanceTransactionList struct {
		ListMeta
		Values []*BalanceTransaction `json:"data"`
	}

	BalanceTransferList struct {
		ListMeta
		Values []*BalanceTransfer `json:"data"`
	}

	BalanceBonusList struct {
		ListMeta
		Values []*BalanceBonus `json:"data"`
	}

	BalanceTransferParams struct {
		User        string                 `json:"user,omitempty"`
		Recipient   string                 `json:"recipient,omitempty"`
		Amount      int64                  `json:"amount"`
		Description string                 `json:"description,omitempty"`
		UserFee     int                    `json:"user_fee,omitempty"`
		Order_no    string                 `json:"order_no,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
	}

	BalanceBonus struct {
		Id                 string                 `json:"id"`
		Object             string                 `json:"object"`
		App                string                 `json:"app"`
		Created            int64                  `json:"created"`
		Livemode           bool                   `json:"livemode"`
		Paid               bool                   `json:"paid"`
		Refunded           bool                   `json:"refunded"`
		Amount             int                    `json:"amount"`
		AmountRefunded     int                    `json:"amount_refunded"`
		TimePaid           int64                  `json:"time_paid"`
		User               string                 `json:"user"`
		Order_no           string                 `json:"order_no"`
		BalanceTransaction string                 `json:"balance_transaction"`
		Description        string                 `json:"description"`
		Metadata           map[string]interface{} `json:"metadata"`
	}

	BalanceBonusParams struct {
		Amount      int                    `json:"amount"`
		User        string                 `json:"user"`
		Order_no    string                 `json:"order_no"`
		Description string                 `json:"description,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
	}

	UserReceipt struct {
		User   string `json:"user"`
		Amount int64  `json:"amount"`
	}
	ReceiptsParams struct {
		Receipts    []UserReceipt `json:"receipts,omitempty"`
		Type        string        `json:"type,omitempty"`
		Description string        `json:"description,omitempty"`
	}
)

type (
	Withdrawal struct {
		Id                  string                 `json:"id"`
		Object              string                 `json:"object"`
		App                 string                 `json:"app"`
		Amount              int64                  `json:"amount"`
		Asset_transaction   string                 `json:"asset_transaction"`
		Balance_transaction string                 `json:"balance_transaction"`
		Channel             string                 `json:"channel"`
		Created             int64                  `json:"created"`
		Description         string                 `json:"description"`
		Extra               map[string]interface{} `json:"extra"`
		Failure_msg         string                 `json:"failure_msg"`
		Fee                 int64                  `json:"fee"`
		Livemode            bool                   `json:"livemode"`
		Metadata            map[string]interface{} `json:"metadata"`
		Operation_url       string                 `json:"operation_url"`
		Order_no            string                 `json:"order_no"`
		Source              string                 `json:"source"`
		Status              string                 `json:"status"`
		Time_canceled       int64                  `json:"time_canceled"`
		Time_succeeded      int64                  `json:"time_succeeded"`
		User                string                 `json:"user"`
		User_fee            int64                  `json:"user_fee"`
		Settle_account      string                 `json:"settle_account"`
	}

	WithdrawalParams struct {
		User           string                 `json:"user,omitempty"`
		Amount         int64                  `json:"amount"`
		Channel        string                 `json:"channel,omitempty"`
		User_fee       int64                  `json:"user_fee"`
		Description    string                 `json:"description,omitempty"`
		Extra          map[string]interface{} `json:"extra,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		Order_no       string                 `json:"order_no,omitempty"`
		Settle_account string                 `json:"settle_account,omitempty"`
	}

	WithdrawalList struct {
		ListMeta
		Values                   []*Withdrawal `json:"data"`
		Total_withdrawals_amount int64         `json:"total_withdrawals_amount"`
	}
)

type (
	BatchWithdrawal struct {
		Id               string                 `json:"id"`
		Object           string                 `json:"object"`
		App              string                 `json:"app"`
		Created          int64                  `json:"created"`
		Livemode         bool                   `json:"livemode"`
		Amount           int64                  `json:"amount"`
		Amount_succeeded int64                  `json:"amount_succeeded"`
		Amount_failed    int64                  `json:"amount_failed"`
		Amount_canceled  int64                  `json:"amount_canceled"`
		Count            int64                  `json:"count"`
		Count_succeeded  int64                  `json:"count_succeeded"`
		Count_failed     int64                  `json:"count_failed"`
		Count_canceled   int64                  `json:"count_canceled"`
		Fee              int64                  `json:"fee"`
		Metadata         map[string]interface{} `json:"metadata"`
		Operation_url    string                 `json:"operation_url"`
		Source           string                 `json:"source"`
		Status           string                 `json:"status"`
		TimeFinished     int64                  `json:"time_finished"`
		User_fee         int64                  `json:"user_fee"`
		Withdrawals      struct {
			ListMeta
			Values []*Withdrawal `json:"data"`
		} `json:"withdrawals"`
	}

	BatchWithdrawalParams struct {
		Withdrawals []string               `json:"withdrawals,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Status      string                 `json:"status,omitempty"`
	}

	BatchWithdrawalList struct {
		ListMeta
		Values []*BatchWithdrawal `json:"data"`
	}
)

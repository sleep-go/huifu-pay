package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2QuickbuckleBankQuery 银行列表查询
//
// 应用场景
// 提供商户、服务商查询快捷、网银、手机WAP、一键绑卡支持的银行列表查询；
//
// 适用对象
// 开通网银、一键绑卡、快捷、手机WAP支付的商户
//
// 接口说明
// 请求URL：https://api.huifu.com/v2/quickbuckle/bank/query
// 请求方式：POST
// 支持格式：JSON
func (t *Trade) V2QuickbuckleBankQuery(req V2QuickbuckleBankQueryRequest) (res *V2QuickbuckleBankQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2QuickbuckleBankQueryRequest(BsPaySdk.V2QuickbuckleBankQueryRequest{
		ReqSeqId:    req.ReqSeqId,
		ReqDate:     req.ReqDate,
		HuifuId:     req.HuifuId,
		BizType:     req.BizType,
		DcType:      req.DcType,
		ExtendInfos: nil,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2QuickbuckleBankQueryResponse](resp)
}

type V2QuickbuckleBankQueryRequest struct {
	ReqSeqId string `json:"req_seq_id"`
	ReqDate  string `json:"req_date"`
	BizType  string `json:"biz_type"`
	DcType   string `json:"dc_type"`
	HuifuId  string `json:"huifu_id"`
}
type V2QuickbuckleBankQueryResponse struct {
	Data struct {
		RespDesc           string           `json:"resp_desc"`
		RespCode           string           `json:"resp_code"`
		OnlineBankInfoList []OnlineBankInfo `json:"online_bank_info_list"`
	} `json:"data"`
	Sign string `json:"sign"`
}
type OnlineBankInfo struct {
	Extend common.StringObject[struct {
		SingleLimitAmt string `json:"singleLimitAmt"`
		DayLimitAmt    string `json:"dayLimitAmt"`
	}] `json:"extend"`
	BankCode         string `json:"bank_code"`
	BankShortEng     string `json:"bank_short_eng"`
	BizType          string `json:"biz_type"`
	BankName         string `json:"bank_name"`
	DcType           string `json:"dc_type"`
	BankAcctIssrId   string `json:"bank_acct_issr_id"`
	BankAcctIssrName string `json:"bank_acct_issr_name"`
}

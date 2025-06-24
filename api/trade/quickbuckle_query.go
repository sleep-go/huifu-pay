package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2QuickbuckleQuery 快捷绑卡查询
//
// 应用场景
// 客户调用该接口，获取持卡人绑卡结果，以及持卡人的cust_id和银行卡token_no用于发起快捷交易。
//
// 适用对象
// 商户需开通代扣支付功能权限，用户已完成代扣绑卡操作。
//
// 接口说明
// 请求URL：https://api.huifu.com/v2/quickbuckle/query
// 请求方式：POST
// 支持格式：JSON
func (t *Trade) V2QuickbuckleQuery(req V2QuickbuckleQueryRequest) (res *V2QuickbuckleQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2QuickbuckleQueryRequest(BsPaySdk.V2QuickbuckleQueryRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		OutCustId:   req.OutCustId,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2QuickbuckleQueryResponse](resp)
}

type V2QuickbuckleQueryExtendInfo struct {
	UserHuifuId string `json:"user_huifu_id"`
}
type V2QuickbuckleQueryRequest struct {
	ReqSeqId    string `json:"req_seq_id"`
	ReqDate     string `json:"req_date"`
	HuifuId     string `json:"huifu_id"`
	OutCustId   string `json:"out_cust_id"`
	ExtendInfos V2QuickbuckleQueryExtendInfo
}
type V2QuickbuckleQueryResponse struct {
	Data struct {
		UserHuifuId string                          `json:"user_huifu_id"`
		RespDesc    string                          `json:"resp_desc"`
		ReqSeqId    string                          `json:"req_seq_id"`
		ReqDate     string                          `json:"req_date"`
		RespCode    string                          `json:"resp_code"`
		HuifuId     string                          `json:"huifu_id"`
		CardList    common.StringObject[[]CardList] `json:"card_list"`
	} `json:"data"`
	Sign string `json:"sign"`
}
type CardList struct {
	CardName string `json:"card_name"`
	CardNo   string `json:"card_no"`
	TokenNo  string `json:"token_no"`
}

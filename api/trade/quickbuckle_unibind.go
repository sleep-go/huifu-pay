package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2QuickbuckleUnbind 新增快捷/代扣解绑接口
//
// 应用场景
// 可调用快捷代扣解绑接口完成快捷绑卡的解约操作。
//
// 适用对象 开通快捷支付商户、服务商；
//
// 接口说明
// 请求URL：https://api.huifu.com/v2/quickbuckle/unbind
// 请求方式：POST
// 支持格式：JSON
func (t *Trade) V2QuickbuckleUnbind(req V2QuickbuckleUnbindRequest) (res *V2QuickbuckleUnbindResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2QuickbuckleUnbindRequest(BsPaySdk.V2QuickbuckleUnbindRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		OutCustId:   req.OutCustId,
		TokenNo:     req.TokenNo,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2QuickbuckleUnbindResponse](resp)
}

type V2QuickbuckleUnbindExtendInfo struct {
	MerchName string `json:"merch_name,omitempty"`
}
type V2QuickbuckleUnbindRequest struct {
	ReqDate     string `json:"req_date,omitempty"`       //请求日期	String	8	Y	日期格式：yyyyMMdd，以北京时间为准；示例值：20220125
	ReqSeqId    string `json:"req_seq_id,omitempty"`     //请求流水号	String	32	Y	同一商户号当天唯一；示例值：2022012614120615001
	HuifuId     string `json:"huifu_id,omitempty"`       //汇付商户Id	String	18	Y	示例值：6666000123123123
	OutCustId   string `json:"org_req_seq_id,omitempty"` //原申请流水号	String	32	Y	和快捷绑卡申请流水号保持一致
	TokenNo     string `json:"org_req_date,omitempty"`   //原申请日期	String	8	Y	和快捷绑卡申请请求日期保持一致
	ExtendInfos V2QuickbuckleUnbindExtendInfo
}
type V2QuickbuckleUnbindResponse struct {
	Data struct {
		RespCode    string `json:"resp_code"`
		RespDesc    string `json:"resp_desc"`
		ReqDate     string `json:"req_date"`
		ReqSeqId    string `json:"req_seq_id"`
		TransStatus string `json:"trans_status"`
	} `json:"data"`
	Sign string `json:"sign"`
}

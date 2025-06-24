package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V3QuickbuckleConfirm 快捷绑卡确认接口
//
// 应用场景
// 持卡人在商户端填写银行短信验证码，完成绑卡确认。收到绑卡成功应答后，商户保存绑卡id。
//
// 适用对象
// 开通快捷支付商户。
//
// 接口说明
// 请求URL：https://api.huifu.com/v3/quickbuckle/confirm
// 请求方式：POST
// 支持格式：JSON
func (t *Trade) V3QuickbuckleConfirm(req V3QuickbuckleConfirmRequest) (res *V3QuickbuckleConfirmResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V3QuickbuckleConfirmRequest(BsPaySdk.V3QuickbuckleConfirmRequest{
		ReqDate:     req.ReqDate,
		ReqSeqId:    req.ReqSeqId,
		HuifuId:     req.HuifuId,
		OrgReqSeqId: req.OrgReqSeqId,
		OrgReqDate:  req.OrgReqDate,
		VerifyCode:  req.VerifyCode,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V3QuickbuckleConfirmResponse](resp)
}

type V3QuickbuckleConfirmRequest struct {
	ReqDate     string `json:"req_date,omitempty"`       //请求日期	String	8	Y	日期格式：yyyyMMdd，以北京时间为准；示例值：20220125
	ReqSeqId    string `json:"req_seq_id,omitempty"`     //请求流水号	String	32	Y	同一商户号当天唯一；示例值：2022012614120615001
	HuifuId     string `json:"huifu_id,omitempty"`       //汇付商户Id	String	18	Y	示例值：6666000123123123
	OrgReqSeqId string `json:"org_req_seq_id,omitempty"` //原申请流水号	String	32	Y	和快捷绑卡申请流水号保持一致
	OrgReqDate  string `json:"org_req_date,omitempty"`   //原申请日期	String	8	Y	和快捷绑卡申请请求日期保持一致
	VerifyCode  string `json:"verify_code,omitempty"`    //验证码	String	6	Y	手机验证码；示例值：345444
}
type V3QuickbuckleConfirmResponse struct {
	Data struct {
		TokenNo      string `json:"token_no"`
		BindCardStat string `json:"bind_card_stat"`
		UserHuifuId  string `json:"user_huifu_id"`
		RespDesc     string `json:"resp_desc"`
		ReqSeqId     string `json:"req_seq_id"`
		ReqDate      string `json:"req_date"`
		RespCode     string `json:"resp_code"`
	} `json:"data"`
	Sign string `json:"sign"`
}

package user

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2UserApplyQuery 用户申请单状态查询
// POST https://api.huifu.com/v2/user/apply/query
//
// 应用场景
// 查询用户非同名对公结算卡审核信息。
//
// 适用对象
// 斗拱服务商或商户；
func (u *User) V2UserApplyQuery(req V2UserApplyQueryRequest) (res *V2UserApplyQueryResponse, raw string, err error) {
	resp, err := u.HuifuPay.BsPay.V2UserApplyQueryRequest(BsPaySdk.V2UserApplyQueryRequest{
		HuifuId:  req.HuifuId,
		ReqSeqId: req.ReqSeqId,
		ReqDate:  req.ReqDate,
		ApplyNo:  req.ApplyNo,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2UserApplyQueryResponse](resp)
}

type V2UserApplyQueryRequest struct {
	ReqSeqId string `json:"req_seq_id"`
	ReqDate  string `json:"req_date"`
	ApplyNo  string `json:"apply_no"`
	HuifuId  string `json:"huifu_id"`
}
type V2UserApplyQueryResponse struct {
	Data struct {
		ApplyReason string `json:"apply_reason"` //审核意见	String	512	N	审核意见
		RespDesc    string `json:"resp_desc"`
		RespCode    string `json:"resp_code"`
		ApplyStatus string `json:"apply_status"` //String	1	N	Y:审核通过 P:审核中 N审核拒绝 F:系统处理失败；示例值：Y
		HuifuId     string `json:"huifu_id"`     //用户号	String	18	N	汇付分配的用户号；示例值：6666000123120000
	} `json:"data"`
	Sign string `json:"sign"`
}

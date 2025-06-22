package user

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2UserBasicdataQuery 用户信息查询
// POST https://api.huifu.com/v2/user/basicdata/query
//
// 应用场景
// 查询为用户基本信息、结算卡及配置费率等信息。
//
// 适用对象 斗拱服务商或商户；
func (u *User) V2UserBasicdataQuery(req V2UserBasicdataQueryRequest) (res *V2UserBasicdataQueryResponse, raw string, err error) {
	resp, err := u.HuifuPay.BsPay.V2UserBasicdataQueryRequest(BsPaySdk.V2UserBasicdataQueryRequest{
		HuifuId:  req.HuifuId,
		ReqSeqId: req.ReqSeqId,
		ReqDate:  req.ReqDate,
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2UserBasicdataQueryResponse](resp)
}

type V2UserBasicdataQueryRequest struct {
	ReqSeqId string `json:"req_seq_id"`
	ReqDate  string `json:"req_date"`
	HuifuId  string `json:"huifu_id"`
}
type V2UserBasicdataQueryResponse struct {
	Data struct {
		RespDesc      string `json:"resp_desc"`
		RespCode      string `json:"resp_code"`
		HuifuId       string `json:"huifu_id"`
		LoginName     string `json:"login_name"`
		LoginPassword string `json:"login_password"`
	} `json:"data"`
	Sign string
}

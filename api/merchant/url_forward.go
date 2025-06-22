package merchant

import (
	"encoding/json"

	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
)

// V2MerchantUrlForward 商户统一进件（页面版）
// 应用场景
// 服务商通过此接口获取商户进件页面地址，商户通过进件页面实现汇付开户、业务开通、微信/支付宝入驻、微信/支付宝实名认证以及与汇付签约等功能，支持商家类型：企业/个体工商户、小微商户。
// 需要注意：一些场景或业务功能的开通需要人工审核，审核通过前即使交易功能开通也会存在交易限额。汇付会尽快完成审核，审核通过后恢复正常交易额度。
//
// 适用对象
// 服务商希望通过汇付页面为商户完成汇付开户以及微信支付宝入驻，简化进件流程的开发。
// 注意：调用接口前渠道商需登录控台配置费率模板。
func (bd *Merchant) V2MerchantUrlForward(req V2MerchantUrlForwardRequest) (res *V2MerchantUrlForwardResponse, raw string, err error) {
	resp, err := bd.HuifuPay.BsPay.V2MerchantUrlForwardRequest(BsPaySdk.V2MerchantUrlForwardRequest{
		ReqSeqId:     req.ReqSeqId,
		ReqDate:      req.ReqDate,
		UpperHuifuId: bd.HuifuPay.BsPay.Msc.SysId,
		StoreId:      req.StoreId,
		ExtendInfos:  BsPaySdk.ToMap(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	marshal, _ := json.Marshal(resp)
	err = json.Unmarshal(marshal, &res)
	if err != nil {
		return nil, "", err
	}
	return res, string(marshal), nil
}

type V2MerchantUrlForwardRequest struct {
	ReqSeqId     string `json:"req_seq_id"`
	ReqDate      string `json:"req_date"`
	UpperHuifuId string `json:"upper_huifu_id"`
	StoreId      string `json:"store_id"` //门店号	String	64	Y	门店号，示例值：SH001
	ExtendInfos  struct {
		Phone           string `json:"phone"`             //手机号	String	11	N	手机号；示例值：13917352618
		Expires         string `json:"expires"`           //跳转地址失效时间	String	5	N	单位毫秒；示例值：50000
		BackPageUrl     string `json:"back_page_url"`     //返回页面URL	String	256	N	进件完成后回跳指定页面。http或https开头； 示例值：http://service.huifu.com/xxx?grantsId=xx&encryptMsg=xxxx
		AsyncReceiveUrl string `json:"async_receive_url"` //异步接收URL	String	256	N	消息接收地址，示例值：http://service.example.com/to/path
		TemplateId      string `json:"template_id"`       //模版编号	String	20	N	服务商控制台，商户入驻费率模版列表中模版编号；示例值：155677
	}
}
type V2MerchantUrlForwardResponse struct {
	Data struct {
		RespDesc  string `json:"resp_desc"`
		ReqSeqId  string `json:"req_seq_id"`
		ProductId string `json:"product_id"`
		ReqDate   string `json:"req_date"`
		RespCode  string `json:"resp_code"`
		Url       string `json:"url"`
	} `json:"data"`
	Sign string `json:"sign"`
}

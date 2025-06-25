package trade

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
	"github.com/sleep-go/huifu-pay/common"
)

// V2TradeSettlementQuery 结算记录查询
// POSThttps://api.huifu.com/v2/merchant/basicdata/settlement/query
//
// 应用场景
// 商户可查询在汇付的所有结算记录及状态。
//
// 适用对象
// 需要确认结算状态以及所有结算记录明细的商户，可调用该功能。查询时间范围100天。
//
// 接口说明
// 支持格式：JSON
func (t *Trade) V2TradeSettlementQuery(req V2TradeSettlementQueryRequest) (res *V2TradeSettlementQueryResponse, raw string, err error) {
	resp, err := t.HuifuPay.BsPay.V2TradeSettlementQueryRequest(BsPaySdk.V2TradeSettlementQueryRequest{
		HuifuId:     req.HuifuId,
		OrgReqDate:  req.OrgReqDate,
		OrgHfSeqId:  req.OrgHfSeqId,
		OrgReqSeqId: req.OrgReqSeqId,
		ExtendInfos: common.StructToMapClean(req.ExtendInfos),
	})
	if err != nil {
		return nil, "", err
	}
	return common.HandleResponse[V2TradeSettlementQueryResponse](resp)
}

type V2TradeSettlementQueryExtendInfo struct {
	ReqSeqId    string `json:"req_seq_id"`
	ReqDate     string `json:"req_date"`
	BeginDate   string `json:"begin_date"`
	EndDate     string `json:"end_date"`
	PageSize    string `json:"page_size"`
	SettleCycle string `json:"settle_cycle"`
	PageNum     string `json:"page_num"`
	TransStat   string `json:"trans_stat"`
	SortColumn  string `json:"sort_column"`
	SortOrder   string `json:"sort_order"`
}
type V2TradeSettlementQueryRequest struct {
	OrgReqDate  string
	OrgHfSeqId  string
	OrgReqSeqId string
	HuifuId     string `json:"huifu_id"`
	ExtendInfos V2TradeSettlementQueryExtendInfo
}
type V2TradeSettlementQueryResponse struct {
	Data struct {
		RespDesc           string `json:"resp_desc"`
		RespCode           string `json:"resp_code"`
		ResultCount        int    `json:"result_count"`
		PageSize           int    `json:"page_size"`
		PageNum            int    `json:"page_num"`
		TransLogResultList []struct {
			TransId        string `json:"trans_id"`
			TransStat      string `json:"trans_stat"`
			HuifuId        string `json:"huifu_id"`
			SettleCycle    string `json:"settle_cycle"`
			CardNo         string `json:"card_no"`
			CardName       string `json:"card_name"`
			BankCode       string `json:"bank_code"`
			TransAmt       string `json:"trans_amt"`
			FeeAmt         string `json:"fee_amt"`
			FeeCustId      string `json:"fee_cust_id"`
			TransDate      string `json:"trans_date"`
			SettleType     string `json:"settle_type"`
			SettleAbstract string `json:"settle_abstract"`
			SettleBatchNo  string `json:"settle_batch_no"`
		} `json:"trans_log_result_list"`
	} `json:"data"`
	Sign string `json:"sign"`
}

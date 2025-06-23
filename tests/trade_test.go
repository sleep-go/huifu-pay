package tests

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/huifurepo/bspay-go-sdk/ut/tool"
	"github.com/sleep-go/huifu-pay/api/trade"
	"github.com/sleep-go/huifu-pay/common"
)

var tr *trade.Trade

func init() {
	huifuPay := common.NewHuifuPay(true, "../config/config.json")
	tr = trade.NewTrade(huifuPay)
}

func TestParseCallback(t *testing.T) {
	http.HandleFunc("/callback", callbackHandler)
	log.Println("监听 8080 中... POST /callback")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	callback, raw, err := common.ParseCallbackRespData[map[string]any](r, tr.HuifuPay.BsPay.Msc)
	if err != nil {
		return
	}
	_ = fileutil.WriteStringToFile("./callback.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", callback)
	// 回复对方
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func TestJSpay(t *testing.T) {
	response, raw, err := tr.V3TradePaymentJspay(trade.V3TradePaymentJspayRequest{
		ReqDate:   tool.GetCurrentDate(),
		ReqSeqId:  tool.GetReqSeqId(),
		HuifuId:   tr.HuifuPay.BsPay.Msc.SysId,
		AcctId:    "",
		GoodsDesc: "GoodsDesc",
		TradeType: common.TradeType_A_NATIVE,
		TransAmt:  "0.01",
		ExtendInfos: trade.V3TradePaymentJspayExtendInfos{
			NotifyUrl: "http://xeab8d5b.natappfree.cc/callback",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}
func TestV3TradePaymentScanpayQuery(t *testing.T) {
	response, raw, err := tr.V3TradePaymentScanpayQuery(trade.V3TradePaymentScanpayQueryRequest{
		HuifuId:     tr.HuifuPay.BsPay.Msc.SysId,
		OrgReqDate:  "20250618",
		OutOrdId:    "",
		OrgHfSeqId:  "0056default250618150144P870ac1367c100000",
		OrgReqSeqId: "RH20250618150110504170",
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}

func TestV2TradePaymentScanpayClose(t *testing.T) {
	response, raw, err := tr.V2TradePaymentScanpayClose(trade.V2TradePaymentScanpayCloseRequest{
		ReqDate:    tool.GetCurrentDate(),
		ReqSeqId:   tool.GetReqSeqId(),
		HuifuId:    tr.HuifuPay.BsPay.Msc.SysId,
		OrgReqDate: "20250623",
		ExtendInfos: trade.V3TradePaymentScanpayCloseExtendInfos{
			OrgHfSeqId:  "",
			OrgReqSeqId: "202506230927021853",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	_ = fileutil.WriteStringToFile("./1.json", raw, false)
	fmt.Printf("======info=======\n%+v\n", response)
}

func TestParseUrl(t *testing.T) {
	ss := `resp_desc=%E4%BA%A4%E6%98%93%E6%88%90%E5%8A%9F%5B000%5D&resp_code=00000000&sign=c85u9v4IrschA6SC9fea4HtymMD3ziOUtwtAZt9%2BO9kci1X4hAfFQQ1TgkfBcHB5M%2BPbflGhd%2F%2F9wGYcASpD0EfhhTMBnGZzNKhQD1%2BnQRryuLlQHtGhDLuB%2FGC2E2JdTE3Xqza9hHWK%2BPdnucoTmVEF2sxpnrxuIyTt0u6SlRDJpVbg7uaBEzrkAOYAMYGu27%2BrcAw2t96VEO9ICSg4IHB2xZS9YceGJjWsGmq26eX0nnlS9uLRUZn%2BUcR902z91Xf5v3S%2BFqMtYdOnWHV3ciCLLhHz%2B9AoYrJW99aEZaOSyDKWhczfKTAJpQuSTWZLJ21UlF1eh8NO1aKxhWApvQ%3D%3D&resp_data=%7B%22acct_date%22%3A%2220250623%22%2C%22acct_id%22%3A%22A44095264%22%2C%22acct_split_bunch%22%3A%7B%22acct_infos%22%3A%5B%7B%22acct_date%22%3A%2220250623%22%2C%22acct_id%22%3A%22A44095264%22%2C%22div_amt%22%3A%220.01%22%2C%22huifu_id%22%3A%226666000168966991%22%7D%5D%2C%22fee_acct_date%22%3A%2220250623%22%2C%22fee_acct_id%22%3A%22A44095264%22%2C%22fee_amt%22%3A%220.00%22%2C%22fee_huifu_id%22%3A%226666000168966991%22%7D%2C%22acct_stat%22%3A%22S%22%2C%22alipay_response%22%3A%7B%22app_id%22%3A%22%22%2C%22buyer_id%22%3A%222088002691234702%22%2C%22buyer_logon_id%22%3A%22357***%2540qq.com%22%2C%22coupon_fee%22%3A%220.00%22%2C%22fund_bill_list%22%3A%5B%7B%22amount%22%3A%220.01%22%2C%22fund_channel%22%3A%22BANKCARD%22%2C%22fund_type%22%3A%22CREDIT_CARD%22%7D%5D%7D%2C%22atu_sub_mer_id%22%3A%222088470620283882%22%2C%22avoid_sms_flag%22%3A%22%22%2C%22bagent_id%22%3A%226666000119490999%22%2C%22bank_desc%22%3A%22TRADE_SUCCESS%22%2C%22bank_message%22%3A%22TRADE_SUCCESS%22%2C%22bank_order_no%22%3A%22202025062322001434701435464536%22%2C%22bank_seq_id%22%3A%22351536%22%2C%22base_acct_id%22%3A%22A44095264%22%2C%22batch_id%22%3A%22250623%22%2C%22channel_type%22%3A%22U%22%2C%22combinedpay_data%22%3A%5B%5D%2C%22combinedpay_fee_amt%22%3A%220.00%22%2C%22debit_flag%22%3A%222%22%2C%22debit_type%22%3A%22C%22%2C%22delay_acct_flag%22%3A%22N%22%2C%22div_flag%22%3A%220%22%2C%22end_time%22%3A%2220250623092828%22%2C%22fee_amount%22%3A%220.00%22%2C%22fee_amt%22%3A%220.00%22%2C%22fee_flag%22%3A2%2C%22fee_formula_infos%22%3A%5B%7B%22fee_formula%22%3A%22MAX%280.00%2CAMT*0.0025%29%22%2C%22fee_type%22%3A%22TRANS_FEE%22%7D%5D%2C%22fee_rec_type%22%3A%221%22%2C%22fee_type%22%3A%22INNER%22%2C%22gate_id%22%3A%22Vu%22%2C%22hf_seq_id%22%3A%22002900TOP4A250623092736P004ac13637000000%22%2C%22huifu_id%22%3A%226666000168966991%22%2C%22is_delay_acct%22%3A%220%22%2C%22is_div%22%3A%220%22%2C%22maze_resp_code%22%3A%22%22%2C%22mer_name%22%3A%22%E5%A4%A9%E6%B4%A5%E7%BA%B3%E5%85%B0%E4%BA%91%E7%A7%91%E6%8A%80%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8%22%2C%22mer_ord_id%22%3A%22202506230927351536%22%2C%22mypaytsf_discount%22%3A%220.00%22%2C%22need_big_object%22%3Afalse%2C%22notify_type%22%3A2%2C%22org_auth_no%22%3A%22%22%2C%22org_huifu_seq_id%22%3A%22%22%2C%22org_trans_date%22%3A%22%22%2C%22out_ord_id%22%3A%22202025062322001434701435464536%22%2C%22out_trans_id%22%3A%22202025062322001434701435464536%22%2C%22party_order_id%22%3A%2203212506233405616309104%22%2C%22pay_amt%22%3A%220.01%22%2C%22pay_scene%22%3A%2202%22%2C%22posp_seq_id%22%3A%2203212506233405616309104%22%2C%22product_id%22%3A%22PAYUN%22%2C%22ref_no%22%3A%22092736351536%22%2C%22req_date%22%3A%2220250623%22%2C%22req_seq_id%22%3A%22202506230927351536%22%2C%22resp_code%22%3A%2200000000%22%2C%22resp_desc%22%3A%22%E4%BA%A4%E6%98%93%E6%88%90%E5%8A%9F%22%2C%22risk_check_data%22%3A%7B%7D%2C%22risk_check_info%22%3A%7B%7D%2C%22settlement_amt%22%3A%220.01%22%2C%22sub_resp_code%22%3A%2200000000%22%2C%22sub_resp_desc%22%3A%22%E4%BA%A4%E6%98%93%E6%88%90%E5%8A%9F%22%2C%22subsidy_stat%22%3A%22I%22%2C%22sys_id%22%3A%226666000168966991%22%2C%22trade_type%22%3A%22A_NATIVE%22%2C%22trans_amt%22%3A%220.01%22%2C%22trans_date%22%3A%2220250623%22%2C%22trans_fee_allowance_info%22%3A%7B%22actual_fee_amt%22%3A%220.00%22%2C%22allowance_fee_amt%22%3A%220.00%22%2C%22allowance_type%22%3A%220%22%2C%22receivable_fee_amt%22%3A%220.00%22%7D%2C%22trans_stat%22%3A%22S%22%2C%22trans_time%22%3A%22092736%22%2C%22trans_type%22%3A%22A_NATIVE%22%2C%22ts_encash_detail%22%3A%5B%5D%7D`
	values, err := url.ParseQuery(ss)
	if err != nil {
		return
	}

	fmt.Println(values)
}

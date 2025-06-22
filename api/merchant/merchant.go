package merchant

import "github.com/sleep-go/huifu-pay/common"

type Merchant struct {
	*common.HuifuPay
}

func NewMerchant(huifuPay *common.HuifuPay) *Merchant {
	return &Merchant{HuifuPay: huifuPay}
}

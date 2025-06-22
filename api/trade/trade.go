package trade

import "github.com/sleep-go/huifu-pay/common"

type Trade struct {
	HuifuPay *common.HuifuPay
}

func NewTrade(huifuPay *common.HuifuPay) *Trade {
	return &Trade{HuifuPay: huifuPay}
}

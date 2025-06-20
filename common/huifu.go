package common

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
)

type HuifuPay struct {
	BsPay *BsPaySdk.BsPay
}

func NewHuifuPay(bsPay *BsPaySdk.BsPay) *HuifuPay {
	return &HuifuPay{BsPay: bsPay}
}

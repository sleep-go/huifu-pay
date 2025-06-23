package common

import (
	"github.com/huifurepo/bspay-go-sdk/BsPaySdk"
)

type HuifuPay struct {
	BsPay *BsPaySdk.BsPay
}

func NewHuifuPay(isProd bool, configPath string) *HuifuPay {
	pay, err := BsPaySdk.NewBsPay(isProd, configPath)
	if err != nil {
		return nil
	}
	return &HuifuPay{BsPay: pay}
}

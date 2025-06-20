package common

import (
	"encoding/json"

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

type StringDecoder[T any] interface {
	Decode() T
}

type StringObject[T any] string

func (o StringObject[T]) Decode() T {
	var res T
	_ = json.Unmarshal([]byte(o), &res)
	return res
}

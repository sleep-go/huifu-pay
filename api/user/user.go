package user

import "github.com/sleep-go/huifu-pay/common"

type User struct {
	HuifuPay *common.HuifuPay
}

func NewUser(huifuPay *common.HuifuPay) *User {
	return &User{HuifuPay: huifuPay}
}

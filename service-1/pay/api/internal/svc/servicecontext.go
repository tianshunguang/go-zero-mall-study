package svc

import (
	"mall/service-1/pay/api/internal/config"
	"mall/service-1/pay/rpc/pay"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	PayRpc pay.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: pay.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}

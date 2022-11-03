package svc

import (
	"mall/service-1/order/api/internal/config"
	"mall/service-1/order/rpc/order"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}

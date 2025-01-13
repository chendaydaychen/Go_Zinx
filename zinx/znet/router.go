package znet

import "Zinx/zinx/ziface"

// 实现router时，先嵌入BaseRouter基类，然后根据需求对基类方法进行重写
type BaseRouter struct{}

// 之所以BaseRouter的方法为空，是因为有的Router不需要PreHandle和PostHandle，
// 所以Router可以继承BaseRouter，不需要实现PreHandle和PostHandle
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

func (br *BaseRouter) Handle(request ziface.IRequest) {}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {}

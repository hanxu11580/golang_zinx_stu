package ziface

type IMsghandler interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgId uint32, router IRouter) //为消息添加具体的处理逻辑
}

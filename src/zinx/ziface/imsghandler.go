package ziface

type IMsghandler interface {
	DoMsgHandler(request IRequest)
	AddRouter(msgId uint32, router IRouter) //为消息添加具体的处理逻辑

	StartWorkerPool()                    //启动worker工作池
	SendMsgToTaskQueue(request IRequest) //将消息交给TaskQueue,由worker进行处理
}

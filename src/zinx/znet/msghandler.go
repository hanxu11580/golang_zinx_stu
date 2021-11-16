package znet

import (
	"fmt"
	"project_stu/src/zinx/utils"
	"project_stu/src/zinx/ziface"
)

// 提供消息路由机制
type MsgHandler struct {
	Apis map[uint32]ziface.IRouter //存放每个MsgId 所对应的处理方法的map属性

	WorkerPoolSize uint32
	TaskQueue      []chan ziface.IRequest //相当于多条管道传输IRequest
}

func NewMsgHandle() *MsgHandler {
	mh := MsgHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
	return &mh
}

func (msgh *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	handler, ok := msgh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("未注册路由")
		return
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}

func (msgh *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := msgh.Apis[msgId]; ok {
		panic("重复添加路由" + fmt.Sprint(msgId))
	}

	msgh.Apis[msgId] = router
	fmt.Println("Add api msgId = ", msgId)
}

// 当多个消息来了，多个go程向消息队列写入
func (msgh *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(msgh.WorkerPoolSize); i++ {
		msgh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		go msgh.WriteToTaskQueue(msgh.TaskQueue[i])
	}
}

func (msgh *MsgHandler) WriteToTaskQueue(taskQueue chan ziface.IRequest) {
	fmt.Println("写入一条请求...")

	// 如果没有就阻塞
	// 只要可以从这条消息队列中 取出数据 就去执行对应的方法，执行完成回来继续循环
	for {
		select {
		case req := <-taskQueue:
			{
				msgh.DoMsgHandler(req)
			}
		}
	}
}

func (msgh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {

	workerIndex := request.GetConnection().GetConnID() % msgh.WorkerPoolSize

	fmt.Println("将消息放入", workerIndex)

	msgh.TaskQueue[workerIndex] <- request
}

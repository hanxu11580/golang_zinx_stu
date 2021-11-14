package znet

import (
	"fmt"
	"project_stu/src/zinx/ziface"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter //存放每个MsgId 所对应的处理方法的map属性
}

func NewMsgHandle() *MsgHandler {
	mh := MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
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

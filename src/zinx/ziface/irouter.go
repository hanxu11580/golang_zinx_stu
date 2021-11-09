package ziface

/*
	从接口函数 可以看出来 IRouter就是处理请求的
	分别为 前 中 后阶段
*/

type IRouter interface {
	PreHandle(request IRequest)  //在处理conn业务之前的钩子方法
	Handle(request IRequest)     //处理conn业务的方法
	PostHandle(request IRequest) //处理conn业务之后的钩子方法
}

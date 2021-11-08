package znet

import (
	"fmt"
	"net"
	"project_stu/src/zinx/ziface"
	"time"
)

//iServer 接口实现，定义一个Server服务类
type Server struct {
	//服务器的名称
	Name string
	//tcp4 or other
	IPVersion string
	//服务绑定的IP地址
	IP string
	//服务绑定的端口
	Port int
}

func (s *Server) Start() {
	fmt.Printf("开始解析服务IP: %s, Port %d\n", s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprint("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("解析TCP地址失败: ", err)
			return
		}

		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("监听失败", s.IPVersion, "err", err)
			return
		}

		fmt.Println("开启服务  ", s.Name, " 成功, 监听中...")

		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept Failed", err)
				continue
			}

			go func() {
				for {
					buf := make([]byte, 512)
					count, err := conn.Read(buf)
					if err != nil {
						fmt.Println("接收失败 ", err)
						continue
					}

					if _, err := conn.Write(buf[:count]); err != nil {
						fmt.Println("写入失败 ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
}

func (s *Server) Serve() {
	s.Start()

	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}

	return s
}

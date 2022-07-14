package main

import (
	"fmt"
	"net"
)

type Server struct {
	Name string
	IP string
	IPVersion string
	Port int
}

func (s *Server) Start(){
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err!=nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")
		for{
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept err, try again ", err)
				continue
			}
			//TODO:最大连接数限制
			//TODO:process方法
			go process(conn)
		}
	}()
	//应该防止start结束，否则go func也结束
	//select和for的区别？
	select {

	}
}

func (s *Server) Serve(){
	s.Start()
	//TODO启动服务的时候还需要处理的事
}
func process(conn net.Conn){

}

func NewServer(name string) *Server{
	return &Server{
		Name :name,
		IPVersion:"tcp4",
		IP:"0.0.0.0",
		Port:7777,
	}
}
func main(){
	s := NewServer("test")
	s.Start()
}
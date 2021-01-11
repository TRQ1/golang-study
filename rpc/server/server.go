package server

import (
		"fmt"
		"log"
		"net"
		"net/rpc"

		"github.com/TRQ1/golang-study/rpc/contract"
)

const port = 1234

type HelloWroldHandler struct{}

func main() {
		log.Printf("Server starting on port %v\n", port)
		StartServer()}

func StartServer() {
		helloWorld := &HelloWroldHandler{}
		rpc.Register(helloWorld)

		l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
				log.Fatal(fmt.Sprintf("Unable to listen on giving port: %s", err))
		}
		defer l.Close().Error()

		for {
		   conn, _ := l.Accept()
		   go rpc.ServeConn(conn)
		}
}

func (h *HelloWroldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
		reply.Message = "Hello" + args.Name
		return nil
}


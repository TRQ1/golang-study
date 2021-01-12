package main

import (
	"fmt"

	"github.com/TRQ1/golang-study/rpc/client"
	"github.com/TRQ1/golang-study/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
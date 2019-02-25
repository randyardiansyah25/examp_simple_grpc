package main

import (
	"fmt"
	"github.com/randyardiansyah25/examp_simple_grpc/testgrpcprotoc/param"
	"github.com/randyardiansyah25/examp_simple_grpc/testgrpcserver/server"
	"google.golang.org/grpc"
	"net"
)

func main() {
	l, _ := net.Listen("tcp", ":9005")
	s := grpc.NewServer()
	h := server.NewServerTest()

	fmt.Println("listen at : 9005")
	param.RegisterTestRequestServer(s, h)
	_ = s.Serve(l)
}

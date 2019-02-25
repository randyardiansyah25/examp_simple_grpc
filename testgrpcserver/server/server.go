package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/randyardiansyah25/examp_simple_grpc/testgrpcprotoc/param"
	"strings"
)

type Server struct{}

func NewServerTest() param.TestRequestServer {
	return &Server{}
}

func (Server) GetResponse(c context.Context, p *param.Request) (*param.Response, error) {
	fmt.Println(">> incoming request: ", p.String())
	if !strings.Contains(p.Email, "@"){
		return nil, errors.New("Invalid email address")
	}

	/*
		uncomment for timeout testing
	 */
	//time.Sleep(60 * time.Second)

	return &param.Response{
		ResponseCode:"0000",
		ResponseMessage:"Success",
	}, nil
}


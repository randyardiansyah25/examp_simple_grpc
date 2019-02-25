package main

import (
	"context"
	"fmt"
	"github.com/kpango/glg"
	"github.com/randyardiansyah25/examp_simple_grpc/testgrpcprotoc/param"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func main() {

	_= glg.Log(">> Dial")
	conn, err := grpc.Dial("localhost:9005", grpc.WithInsecure())
	if err != nil {
		_= glg.Log(">> Finish with error : ", err.Error())
		return
	}

	client := param.NewTestRequestClient(conn)

	req := param.Request{
		Nama:"testing",
		Email:"testing@gmail.com",
	}

	_= glg.Log(">> Call Request")

	/*
		call without deadline
	 */
	//resp, err := client.GetResponse(context.Background(), &req)

	var deadLineValue int64 = 40
	deadline := time.Now().Add(time.Duration(deadLineValue)*time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	resp, err := client.GetResponse(ctx, &req)
	_= glg.Log(">> Finish")
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.Unavailable:
				// refuse or failed handshake with server
				fmt.Println("Unable communicate to server: ", st.Message())
			case codes.DeadlineExceeded:
				fmt.Println("Timeout after",deadLineValue,"second:", st.Message())
			default :
				fmt.Println("Code: ", st.Code(),"; Message : ", st.Message())
			}
		}else{
			fmt.Println("GetResponse Error: ", err.Error())
			return
		}
		return
	}

	fmt.Println("REQ SUCCESS: ", resp.String())
}

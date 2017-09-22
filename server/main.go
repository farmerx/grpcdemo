package main

import (
	example "farmerx/grpcdemo/grpc"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 定义监听地址
const (
	HOST string = "localhost"
	PORT string = "8080"
)

// 定义接口
type FormatData struct{}

func (fd *FormatData) DoFormat(ctx context.Context, in *example.Data) (out *example.Data, err error) {
	str := in.Text
	out = &example.Data{Text: strings.ToUpper(str)}
	return out, nil
}

// 直接在 main 方法中注册接口
func main() {
	listener, err := net.Listen("tcp", HOST+":"+PORT)
	if err != nil {
		log.Fatalln("faile listen at: " + HOST + ":" + PORT)
	} else {
		log.Println("Demo server is listening at: " + HOST + ":" + PORT)
	}
	rpcServer := grpc.NewServer()
	example.RegisterFormatDataServer(rpcServer, &FormatData{})
	reflection.Register(rpcServer)
	if err = rpcServer.Serve(listener); err != nil {
		log.Fatalln("faile serve at: " + HOST + ":" + PORT)
	}
}

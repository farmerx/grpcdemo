package main

import (
	example "farmerx/grpcdemo/grpc"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// 定义请求地址
const (
	ADDRESS string = "localhost:8080"
)

// main 方法实现对 gRPC 接口的请求
func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Can't connect: " + ADDRESS)
	}
	defer conn.Close()
	client := example.NewFormatDataClient(conn)
	resp, err := client.DoFormat(context.Background(), &example.Data{Text: "hello,world!"})
	if err != nil {
		log.Fatalln("Do Format error:" + err.Error())
	}
	log.Println(resp.Text)
}

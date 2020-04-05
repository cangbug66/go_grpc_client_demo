package main

import (
	"context"
	"fmt"
	"go_grpc_client/services"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		log.Fatal("error....", err)
	}
	defer conn.Close()
	c:=services.NewProdServiceClient(conn)
	rsp,err:=c.GetProdStock(context.Background(),&services.ProdRequest{ProdId:1})
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(rsp.ProdStock)
}

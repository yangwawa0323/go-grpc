package main

import (
	"fmt"
	"log"
	"net"

	"51cloudclass.com/go-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	// fmt.Println(services.RegisterProdServiceServer)
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	lis, _ := net.Listen("tcp", ":8080")
	fmt.Println(lis)
	rpcServer.Serve(lis)
}

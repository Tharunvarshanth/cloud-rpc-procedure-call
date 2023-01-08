package main

import (
	"fmt"

	rpchandler "github.com/cloud/rpc/services/server/handler"
	"github.com/cloud/rpc/services/server/proto/pb"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

const (
	//ServiceName to use in discovery
	ServiceName = "server-rpc-service"
)
func main(){
	microService := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(registry.NewRegistry()),
	)

	microService.Init()

	//server := microService.Server()
	// register router
	handler := rpchandler.New() // handlers.NewRPCHandler(ruleEngSvc)

	//Register RPC handlers
	if err := pb.RegisterRpcServerServiceHandler(microService.Server(), handler); err != nil {
		fmt.Printf("unable to register rules engine handler, %v", err)
	}


	if err := microService.Run(); err != nil {
		fmt.Printf("Unable to run the server, %v\n", err)
	}

}


/*
wget https://storage.googleapis.com/golang/go1.19.linux-amd64.tar.gz
tar -xvf golang/go1.19.linux-amd64.tar.gz
mv go go-1.19
sudo mv go-1.19 /usr/local
*/
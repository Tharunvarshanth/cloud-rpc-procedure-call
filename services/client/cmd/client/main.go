package main

import (
	"fmt"

	"github.com/cloud/rpc/services/client/handler"
	"github.com/cloud/rpc/services/client/services"
	"github.com/cloud/rpc/services/server/proto/pb"
	httpServer "github.com/go-micro/plugins/v4/server/http"
	"go-micro.dev/v4"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
)

const (
	SERVER_NAME = "client-http-service" // server name
)

func main() {

	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":8080"),
	)

	httpService := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry.NewRegistry()),
	)
	httpService.Init()
//	c := httpService.Client()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())






    //RPC client
	c := httpService.Client()
	//register rpc services
	serverSvc := pb.NewRpcServerService("server-rpc-service", c)

	clientSrv := services.NewClientService(serverSvc)

	// register router
	handler :=  handler.NewHandler(clientSrv)
	handler.InitRouter(router)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		fmt.Print(err)
	}

	if err := httpService.Run(); err != nil {
		fmt.Printf("Unable to run the server, %v\n", err)
	}

}

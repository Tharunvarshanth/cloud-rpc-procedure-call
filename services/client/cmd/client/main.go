package main

import (
	"fmt"

	"github.com/cloud/rpc/services/client/handler"
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

	// register router
	handler :=  handler.NewHandler()
	handler.InitRouter(router)


	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		fmt.Print(err)
	}

	if err := httpService.Run(); err != nil {
		fmt.Print("Unable to run the server, %v\n", err)
	}

}

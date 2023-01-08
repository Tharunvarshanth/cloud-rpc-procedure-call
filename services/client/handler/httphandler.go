package handler

import (
	"net/http"
	"strconv"

	"github.com/cloud/rpc/services/client/services"
	"github.com/cloud/rpc/services/server/proto/pb"
	"github.com/gin-gonic/gin"
)
type handler struct {
  clientSrv services.ClientService
}

func NewHandler(clientSrv services.ClientService) *handler{
   return &handler{
	clientSrv :clientSrv,
   }
}

func (h *handler) InitRouter(router *gin.Engine) {
	root := router.Group("/api/user/:id")


	root.GET("/",h.getUser )
}

func  (h *handler)  getUser(c *gin.Context) {
    userId, IdErr := strconv.Atoi(c.Param("id"))

	if IdErr != nil ||  userId < 1 {
		c.JSON(http.StatusBadRequest, IdErr)
		return
	}
	res,err:=h.clientSrv.GetAppUserById(c,&pb.GetUserRequest{UserId: c.Param("id")})
    if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if *res == ""{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"name":res,
	})
}

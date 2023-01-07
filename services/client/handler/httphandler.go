package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type handler struct {

}

func NewHandler() *handler{
   return &handler{}
}

func (h *handler) InitRouter(router *gin.Engine) {
	root := router.Group("/api/user")


	root.GET("/",h.getUser )
}

func  (h *handler)  getUser(c *gin.Context) {

	c.JSON(http.StatusOK,gin.H{
		"name":"Tharun",
	})
}

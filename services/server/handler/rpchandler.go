package rpchandler

import (
	"context"
	"fmt"

	"github.com/cloud/rpc/services/server/proto/pb"
)

type handler struct {

}

func New() *handler {
	return &handler{

	}
}


func (h *handler) GetUserById(
	ctx context.Context,
	req *pb.GetUserRequest,
	res *pb.GetUserResponse,

) error {
	fmt.Print("Server Rpc handler")
	if req.UserId != "" {
		fmt.Print("User Id :",req.UserId)
	}
	var users = make(map[string]string)
	users["1"] = "Tharun-1"
	users["2"] = "Tharun-2"
	users["3"] = "Tharun-3"
	users["4"] = "Tharun-4"
	users["5"] = "Tharun-5"

    if users[req.UserId]!=""{
     res.Name=users[req.UserId]
	 return nil
	}
	res.Name = "Not Found"
	return nil

}

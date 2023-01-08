package services

import (
	"context"

	"github.com/cloud/rpc/services/server/proto/pb"
)

type ClientService interface {
	GetAppUserById(context.Context, *pb.GetUserRequest) (*string,error)
}

type clientService struct {
   serverSvc pb.RpcServerService
}

func NewClientService(userSvc pb.RpcServerService) ClientService{
   return &clientService{
    serverSvc: userSvc,
   }
}

func(cs *clientService ) GetAppUserById(ctx context.Context,rq *pb.GetUserRequest)(*string,error){

	user, err := cs.serverSvc.GetUserById(ctx, rq)
	if err != nil {
		return nil, err
	}
	return &user.Name,nil
}

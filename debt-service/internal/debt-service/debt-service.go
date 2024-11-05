package debtservice

import (
	"debt-service/genproto/debtpb"
	load "debt-service/internal/pkg/config"
	"debt-service/internal/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type RunService struct {
	serv *service.Service
}

func NewRunService(serv *service.Service) *RunService {
	return &RunService{
		serv: serv,
	}
}

func (r *RunService) RUN(cfg load.Config) error {
	target := fmt.Sprintf("%s:%d", cfg.DebtServiceHost, cfg.DebtServicePort)
	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	newServer := grpc.NewServer()

	debtpb.RegisterDebtServiceServer(newServer, r.serv)

	return newServer.Serve(listener)
}

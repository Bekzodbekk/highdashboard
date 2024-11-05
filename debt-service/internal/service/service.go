package service

import (
	"context"
	"debt-service/genproto/debtpb"
	"debt-service/internal/postgres"
)

type Service struct {
	debtpb.UnimplementedDebtServiceServer
	repository postgres.DebtRepository
}

func NewService(repos postgres.DebtRepository) *Service {
	return &Service{
		repository: repos,
	}
}

func (s *Service) CreateDebt(ctx context.Context, req *debtpb.CreateDebtReq) (*debtpb.Resp, error) {
	return s.repository.CreateDebt(ctx, req)
}
func (s *Service) UpdateDebt(ctx context.Context, req *debtpb.UpdateDebtReq) (*debtpb.Resp, error) {
	return s.repository.UpdateDebt(ctx, req)
}
func (s *Service) DeleteDebt(ctx context.Context, req *debtpb.DeleteDebtReq) (*debtpb.Resp, error) {
	return s.repository.DeleteDebt(ctx, req)
}
func (s *Service) GetDebtById(ctx context.Context, req *debtpb.GetDebtByIdReq) (*debtpb.Resp, error) {
	return s.repository.GetDebtById(ctx, req)
}
func (s *Service) GetDebtByFilter(ctx context.Context, req *debtpb.GetDebtByFilterReq) (*debtpb.GetDebtByFilterResp, error) {
	return s.repository.GetDebtByFilter(ctx, req)
}

package postgres

import (
	"context"
	"debt-service/genproto/debtpb"
)

type DebtRepository interface {
	CreateDebt(ctx context.Context, req *debtpb.CreateDebtReq) (*debtpb.Resp, error)
	UpdateDebt(ctx context.Context, req *debtpb.UpdateDebtReq) (*debtpb.Resp, error)
	DeleteDebt(ctx context.Context, req *debtpb.DeleteDebtReq) (*debtpb.Resp, error)
	GetDebtById(ctx context.Context, req *debtpb.GetDebtByIdReq) (*debtpb.Resp, error)
	GetDebtByFilter(ctx context.Context, req *debtpb.GetDebtByFilterReq) (*debtpb.GetDebtByFilterResp, error)
}

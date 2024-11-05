package postgres

import (
	"context"
	"database/sql"
	"debt-service/genproto/debtpb"
	"debt-service/storage"

	"github.com/google/uuid"
)

type DebtREPO struct {
	queries *storage.Queries
}

func NewDebtRepo(queries *storage.Queries) DebtRepository {
	return &DebtREPO{
		queries: queries,
	}
}

func (query *DebtREPO) CreateDebt(ctx context.Context, req *debtpb.CreateDebtReq) (*debtpb.Resp, error) {
	resp, err := query.queries.CreateDebt(ctx, storage.CreateDebtParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Jshshir:      int32(req.Jshshir),
		Address:      req.Address,
		Price:        float64(req.Price),
		PricePaid:    float64(req.PricePaid),
		Acquaintance: sql.NullString{String: req.Acquaintance, Valid: req.Acquaintance == ""},
		Collateral:   sql.NullString{String: req.Collateral, Valid: req.Collateral == ""},
		Deadline:     req.Deadline,
	})
	if err != nil {
		return nil, err
	}

	return &debtpb.Resp{
		Status:  true,
		Message: "Create Debt successfuly",
		Debt: &debtpb.Debt{
			Id:           resp.ID.String(),
			FirstName:    resp.FirstName,
			LastName:     resp.LastName,
			PhoneNumber:  resp.PhoneNumber,
			Jshshir:      int64(resp.Jshshir),
			Address:      resp.Address,
			Price:        float32(resp.Price),
			PricePaid:    float32(resp.PricePaid),
			Acquaintance: resp.Acquaintance.String,
			Collateral:   resp.Collateral.String,
			Deadline:     resp.Deadline,
			DebtCUD: &debtpb.DebtCUD{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
				DeletedAt: int64(resp.DeletedAt.Int32),
			},
		},
	}, nil
}
func (query *DebtREPO) UpdateDebt(ctx context.Context, req *debtpb.UpdateDebtReq) (*debtpb.Resp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	resp, err := query.queries.UpdateDebt(ctx, storage.UpdateDebtParams{
		ID:           id,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		PhoneNumber:  req.PhoneNumber,
		Jshshir:      int32(req.Jshshir),
		Address:      req.Address,
		Price:        float64(req.Price),
		PricePaid:    float64(req.PricePaid),
		Acquaintance: sql.NullString{String: req.Acquaintance, Valid: req.Acquaintance == ""},
		Collateral:   sql.NullString{String: req.Collateral, Valid: req.Collateral == ""},
		Deadline:     req.Deadline,
	})
	if err != nil {
		return nil, err
	}

	return &debtpb.Resp{
		Status:  true,
		Message: "Update Debt successfuly",
		Debt: &debtpb.Debt{
			Id:           resp.ID.String(),
			FirstName:    resp.FirstName,
			LastName:     resp.LastName,
			PhoneNumber:  resp.PhoneNumber,
			Jshshir:      int64(resp.Jshshir),
			Address:      resp.Address,
			Price:        float32(resp.Price),
			PricePaid:    float32(resp.PricePaid),
			Acquaintance: resp.Acquaintance.String,
			Collateral:   resp.Collateral.String,
			Deadline:     resp.Deadline,
			DebtCUD: &debtpb.DebtCUD{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
				DeletedAt: int64(resp.DeletedAt.Int32),
			},
		},
	}, nil
}
func (query *DebtREPO) DeleteDebt(ctx context.Context, req *debtpb.DeleteDebtReq) (*debtpb.Resp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	err = query.queries.DeleteDebt(ctx, storage.DeleteDebtParams{
		ID: id,
	})
	if err != nil {
		return nil, err
	}

	return &debtpb.Resp{
		Status:  true,
		Message: "Delete Debt Successfuly",
	}, nil
}

func (query *DebtREPO) GetDebtById(ctx context.Context, req *debtpb.GetDebtByIdReq) (*debtpb.Resp, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	resp, err := query.queries.GetDebtById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &debtpb.Resp{
		Status:  true,
		Message: "Get Debt Successfuly",
		Debt: &debtpb.Debt{
			Id:           resp.ID.String(),
			FirstName:    resp.FirstName,
			LastName:     resp.LastName,
			PhoneNumber:  resp.PhoneNumber,
			Jshshir:      int64(resp.Jshshir),
			Address:      resp.Address,
			Price:        float32(resp.Price),
			PricePaid:    float32(resp.PricePaid),
			Acquaintance: resp.Acquaintance.String,
			Collateral:   resp.Collateral.String,
			Deadline:     resp.Deadline,
			DebtCUD: &debtpb.DebtCUD{
				CreatedAt: resp.CreatedAt.Time.String(),
				UpdatedAt: resp.UpdatedAt.Time.String(),
				DeletedAt: int64(resp.DeletedAt.Int32),
			},
		},
	}, nil
}
func (query *DebtREPO) GetDebtByFilter(ctx context.Context, req *debtpb.GetDebtByFilterReq) (*debtpb.GetDebtByFilterResp, error) {
	resp, err := query.queries.GetDebtByFilter(ctx, req.String())
	if err != nil {
		return nil, err
	}

	debts := make([]*debtpb.Debt, 0, len(resp))
	for _, debt := range resp {
		debts = append(debts, &debtpb.Debt{
			Id:           debt.ID.String(),
			FirstName:    debt.FirstName,
			LastName:     debt.LastName,
			PhoneNumber:  debt.PhoneNumber,
			Jshshir:      int64(debt.Jshshir),
			Address:      debt.Address,
			Price:        float32(debt.Price),
			PricePaid:    float32(debt.PricePaid),
			Acquaintance: debt.Acquaintance.String,
			Collateral:   debt.Collateral.String,
			Deadline:     debt.Deadline,
			DebtCUD: &debtpb.DebtCUD{
				CreatedAt: debt.CreatedAt.Time.String(),
				UpdatedAt: debt.UpdatedAt.Time.String(),
				DeletedAt: int64(debt.DeletedAt.Int32),
			},
		})
	}

	return &debtpb.GetDebtByFilterResp{
		Status: true,
		Message: "Debt Get By Filter Successfuly",
		Debts: debts,
	}, nil
}

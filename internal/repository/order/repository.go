package order

import (
	"context"

	"github.com/xxvlrapss/go_restorant_app.git/internal/model"
)


type Repository interface {
	CreateOrder(ctx context.Context, order model.Order) (model.Order, error)
	GetOrderInfo(ctx context.Context, orderID string) (model.Order, error)
}
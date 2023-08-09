package order

import (
	"github.com/xxvlrapss/go_restorant_app.git/internal/model"
)

type Repository interface { 
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderInfo(orderID string) (model.Order, error)
}
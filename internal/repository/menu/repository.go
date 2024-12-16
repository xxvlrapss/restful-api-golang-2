package menu

import (
	"context"

	"github.com/xxvlrapss/go_restorant_app.git/internal/model"
)

type Repository interface {
	GetMenuList(ctx context.Context, menuType string) ([]model.MenuItem, error)
	GetMenu(ctx context.Context, orderCode string) (model.MenuItem, error)
}
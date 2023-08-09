package menu

import (
	"github.com/xxvlrapss/go_restorant_app.git/internal/model"
)

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}
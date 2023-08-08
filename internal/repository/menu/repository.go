package menu

import "github.com/xxvlrapss/go_restorant_app.git/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
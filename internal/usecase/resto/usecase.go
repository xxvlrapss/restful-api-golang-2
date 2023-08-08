package resto

import "github.com/xxvlrapss/go_restorant_app.git/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
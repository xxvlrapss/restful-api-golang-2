package resto

import (
	"github.com/xxvlrapss/go_restorant_app.git/internal/model"
	"github.com/xxvlrapss/go_restorant_app.git/internal/repository/menu"
)

type restoUsecase struct {
	menuRepo menu.Repository
}

func GetUsecase(menuRepo menu.Repository) Usecase {
	return &restoUsecase{
		menuRepo: menuRepo,
	}
}

func (r *restoUsecase) GetMenu(menuType string) ([]model.MenuItem, error){
	return r.menuRepo.GetMenu(menuType)
} 
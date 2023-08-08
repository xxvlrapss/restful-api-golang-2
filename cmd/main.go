package main

import (
	"github.com/labstack/echo"

	"github.com/xxvlrapss/go_restorant_app.git/internal/database"
	"github.com/xxvlrapss/go_restorant_app.git/internal/delivery/rest"
	mRepo "github.com/xxvlrapss/go_restorant_app.git/internal/repository/menu"
	rUsecase "github.com/xxvlrapss/go_restorant_app.git/internal/usecase/resto"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=leon123 dbname=go_resto_app sslmode=disable"
)

func main() {
	e := echo.New()
	// localhost:14040/menu/food
	db := database.GetDB(dbAddress)

	menuRepo := mRepo.GetRepository(db)

	restoUsecase := rUsecase.GetUsecase(menuRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)

	
 	e.Logger.Fatal(e.Start(":14045"))
}







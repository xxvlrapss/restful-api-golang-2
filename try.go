// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/labstack/echo"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// const (dbAddress = "host=localhost port=5432 user=postgres password=leon12345 dbname=go_resto_app sslmode=disable")

// func main() {
// 	seedDB()
// 	e := echo.New()
// 	// localhost:14040/menu/food
//  e.GET("/menu/food", getMenu)
//  e.GET("/menu/drinks", getMenu)
//  e.Logger.Fatal(e.Start(":14040"))
// }


// type MenuType string

// const (
// 	MenuTypeFood = "food"
// 	MenuTypeDrink = "drink"
// )


// type MenuItem struct { 
// 	Name string
// 	OrderCode string
// 	Price int64
// 	Type MenuType
// }

// type DrinkItem struct { 
// 	Name string
// 	OrderCode string
// 	Price int
// 	Type MenuType
// }


// func seedDB() {
// 	foodMenu := []MenuItem{
// 		{
// 			Name: "Indomie",
// 			OrderCode: "indomie",
// 			Price: 30000,
// 			Type: MenuTypeFood,
// 		},
// 		{
// 			Name: "Ayam lada hitam",
// 			OrderCode: "ayam_lada_hitam",
// 			Price: 40000,
// 			Type: MenuTypeFood,
// 		},
// 		{
// 			Name: "Ayam goreng",
// 			OrderCode: "ayam_goreng",
// 			Price: 40000,
// 			Type: MenuTypeFood,
// 		},
// 	}
// 	drinksMenu := []DrinkItem{
// 		{
// 			Name: "Teh Manis",
// 			OrderCode: "teh_manis",
// 			Price: 12000,
// 			Type: MenuTypeDrink,
// 		},
// 		{
// 			Name: "Jus Jeruk",
// 			OrderCode: "jus_jeruk",
// 			Price: 15000,
// 			Type: MenuTypeDrink,
// 		},
// 	}

// 	fmt.Println(foodMenu, drinksMenu)

	
// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	// db.AutoMigrate(&MenuItem{})

// 	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
// 	db.Create(&foodMenu)
// 	db.Create(&drinksMenu)
// 	}
// }

// func getMenu(c echo.Context) error {
// 	menuType := c.FormValue("menu_type")
	
// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic(err)
// 	}
// 	var menuData []MenuItem

// 	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
// }

// // // func getFoodMenu(c echo.Context) error {
// // // 	db, err := gorm.Open(postgres.Open(dbAddress))
// // // 	if err != nil {
// // // 		panic(err)
// // // 	}

// // // 	var menuData []MenuItem

// // // 	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)

// // // 	return c.JSON(http.StatusOK, map[string]interface{}{
// // // 		"data": menuData,
// // // 	})
// // // }

// // func getDrinkMenu(c echo.Context) error {
// // 	db, err := gorm.Open(postgres.Open(dbAddress))
// // 	if err != nil {
// // 		panic(err)
// // 	}

// // 	var menuData []MenuItem

// // 	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)
// // 	return c.JSON(http.StatusOK, map[string]interface{}{
// // 		"data": menuData,
// // 	})
// // }






// func getFoodMenu(c echo.Context) error {
// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic(err)
// 	}

// 	var menuData []MenuItem
// 	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)

// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
// }

// func getDrinksMenu(c echo.Context) error {
// 	db, err := gorm.Open(postgres.Open(dbAddress))
// 	if err != nil {
// 		panic(err)
// 	}

// 	var menuData []MenuItem
// 	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)
// 	return c.JSON(http.StatusOK, map[string]interface{}{
// 		"data": menuData,
// 	})
//
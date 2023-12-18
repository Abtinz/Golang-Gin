package main

/* List of urls
GET /basket/ (returns a list of baskets)
POST /basket/ (creates a new basket)
PATCH /basket/<id> (updates the given basket)
GET /basket/<id> (reutrns the given basket)
DELETE /basket/<id> (deletes the given backset)
*/

import (
	"log"
	"net/http"
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/database"
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/model"
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/routers/basket"
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/routers/user"
	"github.com/labstack/echo/v4"
)

func databaseInitializer() {

	//initializering and opening postgress database
	database.InitDB()
	defer database.CloseDB()
	//now is time to migrate our models to database
	database.DB.AutoMigrate(&model.Basket{})
	database.DB.AutoMigrate(&model.User{})
}
func main() {

	databaseInitializer()

	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "welcome to your basket!")
	})

	BASE_URL := e.Group("/basket")
	USER_BASE_URL := e.Group("/user")
	host := "0.0.0.0:8080" //no where is like home ...

	BASE_URL.GET("/p", func(context echo.Context) error {
		return context.String(http.StatusOK, "Ping!")
	})

	BASE_URL.GET("/", basket.AllBaskets)
	BASE_URL.POST("/", basket.CreateBasket)
	BASE_URL.GET("/:id", basket.GetBasketByID)
	BASE_URL.PATCH("/:id", basket.UpdateBasket)
	BASE_URL.DELETE("/:id", basket.DeleteBasket)
	USER_BASE_URL.GET("/:id", user.GetUserByID)
	USER_BASE_URL.POST("/signup/", user.SignUp)
	USER_BASE_URL.POST("/login/", user.LoginUser)
	USER_BASE_URL.DELETE("/delete/:id", user.RemoveTheAccount)

	log.Fatal(e.Start(host))
}

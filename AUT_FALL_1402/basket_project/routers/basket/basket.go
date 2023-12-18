package basket

import (

	//self libraries
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/database"
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/model"
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/routers/authentication"

	//OTHERS
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// this function will create our user baskets in json response
func CreateBasket(c echo.Context) error {

	userID, err := authentication.SolveTokenUserIDFuntion(c)
	if err != nil {
		//user is unauthorized we have to echo this error for all user security, token is invalid or ...
		return echo.ErrUnauthorized
	}

	basket := new(model.Basket)
	basket.UserID = userID
	if err := c.Bind(basket); err != nil {
		//data is not found or user token expired or ...
		//the server cannot or will not process the request due to something that is perceived to be a client error ...
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusBadRequest, message)
	}

	//lets create user basket and see what will be happen next
	if err := database.DB.Create(&basket).Error; err != nil {
		//500 error, check the logs for data base error something is wrong fro, our processes
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}

	//every thing is ok and we will return 200 status code with user baskets
	return c.JSON(http.StatusCreated, basket)
}

// this function will load our user baskets in json response
func AllBaskets(c echo.Context) error {

	_, err := authentication.SolveTokenUserIDFuntion(c)
	if err != nil {
		//user is unauthorized we have to echo this error for all user security, token is invalid or ...
		return echo.ErrUnauthorized
	}

	//baskets will be used in returned
	var baskets []model.Basket
	if err := database.DB.Find(&baskets).Error; err != nil {
		//500 error, check the logs for data base error something is wrong fro, our processes
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}

	//every thing is ok and we will return 200 status code with user baskets
	return c.JSON(http.StatusOK, baskets)
}

// update your basket
func UpdateBasket(c echo.Context) error {

	_, err := authentication.SolveTokenUserIDFuntion(c)
	if err != nil {
		//user is unauthorized we have to echo this error for all user security, token is invalid or ...
		return echo.ErrUnauthorized
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//user inputed id for updating is wrong ... 400 status code
		message := map[string]string{"error_message": "wrong id try again later"}
		return c.JSON(http.StatusBadRequest, message)
	}

	var basket model.Basket
	if err := database.DB.First(&basket, id).Error; err != nil {
		//404, well know error for basket search result ...
		message := map[string]string{"error_message": "im so sorry but this id is not related to any basket"}
		return c.JSON(http.StatusNotFound, message)
	}

	if err := c.Bind(&basket); err != nil {
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusBadRequest, message)
	}

	if err := database.DB.Save(&basket).Error; err != nil {
		//500 error, check the logs for data base error something is wrong fro, our processes
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}

	//we found user basket with its special id!
	return c.JSON(http.StatusOK, basket)
}

func GetBasketByID(c echo.Context) error {
	_, err := authentication.SolveTokenUserIDFuntion(c)
	if err != nil {
		//user is unauthorized we have to echo this error for all user security, token is invalid or ...
		return echo.ErrUnauthorized
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//user inputed id for updating is wrong ... 400 status code
		message := map[string]string{"error_message": "wrong id try again later"}
		return c.JSON(http.StatusBadRequest, message)
	}

	var basket model.Basket
	if err := database.DB.First(&basket, id).Error; err != nil {
		//404, well know error for basket search result ...
		message := map[string]string{"error_message": "im so sorry but this id is not related to any basket"}
		return c.JSON(http.StatusNotFound, message)
	}

	return c.JSON(http.StatusOK, basket)

}

func DeleteBasket(c echo.Context) error {
	_, err := authentication.SolveTokenUserIDFuntion(c)
	if err != nil {
		//user is unauthorized we have to echo this error for all user security, token is invalid or ...
		return echo.ErrUnauthorized
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//user inputted id for updating is wrong ... 400 status code
		message := map[string]string{"error_message": "wrong id try again later"}
		return c.JSON(http.StatusBadRequest, message)
	}

	var basket model.Basket
	if err := database.DB.First(&basket, id).Error; err != nil {
		//404, well know error for basket search result ...
		message := map[string]string{"error_message": "im so sorry but this id is not related to any basket"}
		return c.JSON(http.StatusNotFound, message)
	}

	if err := database.DB.Delete(&basket).Error; err != nil {
		//500 error, check the logs for data base error something is wrong fro, our processes
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}

	message := map[string]string{"message": "Basket have been Deleted from basket Successfully"}
	return c.JSON(http.StatusOK, message)

}

package user

import (
	//self packages imports
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/database"
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/model"
	"github.com/bradtraversy/AUT_FALL_1402/basket_project/routers/authentication"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// we will need this for resolving user json request for authorization
type LoginInfoJson struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// THIS FUNCTION WILL MAKE A NEW USER IN DATA BASE
func SignUp(c echo.Context) error {

	//new user is created for binding
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		//data is not found or ...
		//the server cannot or will not process the request due to something that is perceived to be a client error ...
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusBadRequest, message)
	}

	if err := database.DB.Create(&user).Error; err != nil {
		//500 error, check the logs for data base error something is wrong fro, our processes
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}

	//user is created successfully
	return c.JSON(http.StatusCreated, user)
}

// give your id in params and this function will give your the user
func GetUserByID(c echo.Context) error {
	//resolving id from parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//the server cannot or will not process the request due to something that is perceived to be a client error ...
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusBadRequest, message)
	}

	//lets find user from db
	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		//404, well know error for basket search result ...
		message := map[string]string{"error_message": "im so sorry but this id is not related to any basket"}
		return c.JSON(http.StatusNotFound, message)
	}

	//here is identified user response ...
	return c.JSON(http.StatusOK, user)
}

// user and pass is needed for login in json mode
func LoginUser(c echo.Context) error {

	//we will need this for resolving user json request for authorization
	type LoginInfoJson struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	login := new(LoginInfoJson)

	//bad request checking
	if err := c.Bind(login); err != nil {
		//the server cannot or will not process the request due to something that is perceived to be a client error ...
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusBadRequest, message)
	}

	var user model.User
	if err := database.DB.Where("username = ?", login.Username).First(&user).Error; err != nil {
		//username is invalid 401 is coming ...
		message := map[string]string{"error_message": "Invalid username .."}
		return c.JSON(http.StatusUnauthorized, message)
	}

	if user.Password != login.Password {
		//password is invalid 401 is coming ...
		message := map[string]string{"error_message": "Invalid password ... (username is correct)"}
		return c.JSON(http.StatusUnauthorized, message)
	}

	token, err := authentication.GenerateToken(user.UserID)
	if err != nil {
		//500 error, check the logs for data base error something is wrong fro, our processes
		//in this case i think its about our token generator!
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}

	generated_token := map[string]string{"token": token}
	return c.JSON(http.StatusOK, generated_token)
}

// this function will remove account(id is needed)
func RemoveTheAccount(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		//the server cannot or will not process the request due to something that is perceived to be a client error ...
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusBadRequest, message)
	}

	var user model.User
	if err := database.DB.First(&user, id).Error; err != nil {
		//404, well know error for basket search result ...
		message := map[string]string{"error_message": "im so sorry but this id is not related to any basket"}
		return c.JSON(http.StatusNotFound, message)
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		//500 error, check the logs for data base error something is wrong fro, our processes
		message := map[string]string{"error_message": err.Error()}
		return c.JSON(http.StatusInternalServerError, message)
	}

	//bye bye user ...
	return c.NoContent(http.StatusNoContent)
}

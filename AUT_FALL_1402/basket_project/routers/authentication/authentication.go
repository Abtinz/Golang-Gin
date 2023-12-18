package authentication

import (
	"net/http"
	"strings"
	"time"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

//private key, we will sign our tokens with this in generator
var jwtKey = []byte("263526715361253126483719834248098")

func GenerateToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	//lets sign this token (i know it not kerbrose but its enough for localserver trust me ...)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token signing method")
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	return token, nil
}

func ExtractToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "No JWT token found")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization format")
	}

	return parts[1], nil
}

/* this function will return user uuid which is decoded in token with address and others
*/
func SolveTokenUserIDFuntion(c echo.Context) (uint, error) {
	tokenString, err := ExtractToken(c)
	if err != nil {
		return 0, err
	}

	token, err := ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid user ID in token claims")
	}

	userID := uint(userIDFloat)
	return userID, nil
}

package v1

import (
	"net/http"
	"time"

	"framework/domain"
	"framework/repo"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("your_secret_key")

type LoginHandler struct {
	UserRepo *repo.UserRepo
}

func NewLoginHandler(e *echo.Echo, userRepo *repo.UserRepo) {
	h := &LoginHandler{UserRepo: userRepo}

	group := e.Group("/api/v1")
	group.POST("/login", h.Login)
}

func (h *LoginHandler) Login(c echo.Context) error {
	var req domain.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	user, err := h.UserRepo.FindByUsername(req.Username)
	if err != nil || user.Password != req.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}
	token, err := generateJWT(user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "token error"})
	}
	return c.JSON(http.StatusOK, domain.LoginResponse{Token: token})
}

func generateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(jwtSecret)
}

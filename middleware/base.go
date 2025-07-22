package middleware

import "github.com/labstack/echo/v4"

type BaseMiddleware struct{}

func NewBaseMiddleware() *BaseMiddleware {
	return &BaseMiddleware{}
}

func (m *BaseMiddleware) BaseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 通用中间件逻辑
		return next(c)
	}
}

package http

import (
	"frete-rapido/internal/quote"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group, uc quote.UseCase) {
	handler := NewHandler(uc)
	g.POST("/quote", handler.CreateQuote)
	g.GET("/metrics", handler.GetMetrics)
}

package http

import (
	"frete-rapido/internal/entity"
	"frete-rapido/internal/quote"
	"strconv"

	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	QuoteUC quote.UseCase
}

func NewHandler(uc quote.UseCase) *Handler {
	return &Handler{QuoteUC: uc}
}

// POST /api/quote
func (h *Handler) CreateQuote(c echo.Context) error {
	var req entity.QuoteRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "payload inválido"})
	}

	resp, err := h.QuoteUC.Cotar(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, resp)
}

// http/handler.go
func (h *Handler) GetMetrics(c echo.Context) error {
	lastQuotes := 0
	if param := c.QueryParam("last_quotes"); param != "" {
		if n, err := strconv.Atoi(param); err == nil {
			lastQuotes = n
		}
	}

	metrics, err := h.QuoteUC.GetMetrics(c.Request().Context(), lastQuotes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, metrics)
}

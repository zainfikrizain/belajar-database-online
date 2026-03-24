package handler

import (
	"Tugas-2/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	products, err := h.service.GetAllProducts(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to fetch products",
		})
	}

	return c.JSON(http.StatusOK, products)
}

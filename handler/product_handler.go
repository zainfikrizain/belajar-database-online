package handler

import (
	"Tugas-2/models"
	"Tugas-2/services"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	products, err := h.service.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(500, map[string]string{
			"error": "blablabla",
		})
	}

	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) Create(c echo.Context) error {
	var product models.Product

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(), // 👈 see real errorx
		})
	}

	if err := h.service.Create(c.Request().Context(), &product); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to create product",
		})
	}

	return c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "product not found",
			})
		}
		return c.JSON(500, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	product.ID = id

	if err := h.service.Update(&product); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "product not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to update product",
		})
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid id",
		})
	}

	if err := h.service.Delete(id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "product not found",
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to delete product",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "deleted successfully",
	})
}

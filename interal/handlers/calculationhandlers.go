package handlers

import (
	calculationserver "Calculator/interal/calculationServer"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalculationHandlers struct {
	service calculationserver.CalculationService
}

func NewCalculationHandler(s calculationserver.CalculationService) *CalculationHandlers {
	return &CalculationHandlers{service: s}
}

func (h *CalculationHandlers) GetCalculation(c echo.Context) error {
	calculations, err := h.service.GetAllCalculation()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not get calculations"})
	}
	return c.JSON(http.StatusOK, calculations)
}

func (h *CalculationHandlers) PostCalculation(c echo.Context) error {
	var req calculationserver.CalculationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalidaaaaa type"})
	}

	cacl, err := h.service.CreateCalculation(req.Expresion)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create calculation"})
	}

	return c.JSON(http.StatusCreated, cacl)
}

func (h *CalculationHandlers) PatchCalcuator(c echo.Context) error {
	var req calculationserver.CalculationRequest
	id := c.Param("id")

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalyd request"})
	}

	update, err := h.service.UpdateCalculation(id, req.Expresion)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not udate"})
	}
	return c.JSON(http.StatusOK, update)
}

func (h *CalculationHandlers) DeleteCalculator(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCalculetion(id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not delete colculation"})
	}

	return c.NoContent(http.StatusNoContent)
}

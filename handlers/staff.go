package handlers

import (
	"golang-echo/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handlerStaff struct {
	StaffRepository repositories.Staff
}

func HandlerStaff(StaffRepository repositories.Staff) *handlerStaff {
	return &handlerStaff{StaffRepository}
}

func (h *HandlerGlobal) AllStaff(c echo.Context) error {

	returnvalue, _ := h.StaffRepository.GetAllStaff()
	return c.JSON(http.StatusOK, returnvalue)
}

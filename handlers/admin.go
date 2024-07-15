package handlers

import (
	"golang-echo/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handlerAdmin struct {
	StaffRepository repositories.Staff
}

func HandlerAdmin(AdminRepository repositories.Staff) *handlerAdmin {
	return &handlerAdmin{AdminRepository}
}

func (h *HandlerGlobal) AllAdmin(c echo.Context) error {

	returnvalue, _ := h.AdminRepository.GetAllAdmin()
	return c.JSON(http.StatusOK, returnvalue)
}

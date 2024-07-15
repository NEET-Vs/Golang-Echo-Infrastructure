package handlers

import "golang-echo/repositories"

type HandlerGlobal struct {
	StaffRepository repositories.Staff
	AdminRepository repositories.Admin
}

func Handler(handler HandlerGlobal) *HandlerGlobal {
	return &HandlerGlobal{

		StaffRepository: handler.StaffRepository,
		AdminRepository: handler.AdminRepository,
	}
}

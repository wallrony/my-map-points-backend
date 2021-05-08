package handlers

import (
	"my_map_points/core/utils"
	repo_impl "my_map_points/data/repositories/implementation"
	services_abs "my_map_points/services/abstraction"
	services_impl "my_map_points/services/implementation"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service services_abs.UserAbstractService
}

func (handler UserHandler) FindAll(context echo.Context) error {
	res := context.Response()

	result, err := handler.service.FindAll()

	if err.Err != "" {
		utils.WriteResponse(res, 500, err.AsMessage())

		return nil
	}

	utils.WriteResponse(res, 200, result)

	return nil
}

func NewUserHandler() *UserHandler {
	r := repo_impl.NewUserRepository()
	s := services_impl.NewUserService(r)
	h := &UserHandler{s}

	return h
}

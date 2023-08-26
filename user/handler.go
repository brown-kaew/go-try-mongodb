package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	FindById() echo.HandlerFunc
}

type handler struct {
	userDb UserDb
}

func NewHandler(userDb UserDb) Handler {
	return &handler{
		userDb: userDb,
	}
}

func (handler *handler) FindById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid Id")
		}

		user, err := handler.userDb.FindById(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}

		return c.JSON(http.StatusOK, user)
	}
}

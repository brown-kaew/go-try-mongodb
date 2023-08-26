package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	FindById() echo.HandlerFunc
}

type handler struct {
	userDb      UserDb
	simpleRedis SimpleRedis
}

func NewHandler(userDb UserDb, simpleRedis SimpleRedis) Handler {
	return &handler{
		userDb:      userDb,
		simpleRedis: simpleRedis,
	}
}

func (handler *handler) FindById() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId := c.Param("id")
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid Id")
		}

		user := handler.simpleRedis.Get(userId)
		if user != nil {
			return c.JSON(http.StatusOK, user)
		}
		user, err = handler.userDb.FindById(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		if user != nil {
			log.Printf("put redis userId=%s", userId)
			handler.simpleRedis.Put(userId, user)
		}

		return c.JSON(http.StatusOK, user)
	}
}

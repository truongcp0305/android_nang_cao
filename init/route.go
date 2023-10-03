package init

import (
	"android-service/adapter/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, tc controller.TaskController, uc controller.UserController, wc controller.WordController) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "oke")
	})
	e.GET("/task/:id", func(c echo.Context) error {
		return tc.Detail(c)
	})
	e.GET("/task/list-task/:userId", func(c echo.Context) error {
		return tc.GetList(c)
	})
	e.POST("/task", func(c echo.Context) error {
		return tc.Create(c)
	})
	e.PUT("/task", func(c echo.Context) error {
		return tc.Update(c)
	})
	e.DELETE("/task/:id", func(c echo.Context) error {
		return tc.Delete(c)
	})
	e.POST("/user", func(c echo.Context) error {
		return uc.Create(c)
	})
	e.POST("/user/login", func(c echo.Context) error {
		return uc.Login(c)
	})
	e.POST("/word", func(c echo.Context) error {
		return wc.Insert(c)
	})
}

package init

import (
	"android-service/adapter/controller"
	"android-service/infrastructure"
	"log"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, tc controller.TaskController, uc controller.UserController, wc controller.WordController, ws controller.SocketController) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := infrastructure.GetMiddleWareConfig()
	e.Use(echojwt.WithConfig(config))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "oke")
	})
	e.POST("/reset", func(c echo.Context) error {
		return uc.ResetPass(c)
	})
	e.GET("/reset/:encrypt", func(c echo.Context) error {
		return uc.Reset(c)
	})
	e.POST("/user/login", func(c echo.Context) error {
		return uc.Login(c)
	})
	e.GET("/task/all", func(c echo.Context) error {
		return tc.GetAll(c)
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
	e.POST("/user/update", func(c echo.Context) error {
		return uc.Update(c)
	})
	e.GET("/user/get-list", func(c echo.Context) error {
		return uc.GetList(c)
	})
	e.POST("/word", func(c echo.Context) error {
		return wc.Insert(c)
	})
	e.GET("/word/question/:level", func(c echo.Context) error {
		return wc.GetQuestions(c)
	})
	e.POST("/match/status", func(c echo.Context) error {
		return ws.Status(c)
	})
	e.GET("/match/join/:id/:level", func(c echo.Context) error {
		return ws.Join(c)
	})
	e.GET("/match/leave/:id", func(c echo.Context) error {
		return ws.Leave(c)
	})
	e.POST("/task/assign-to", func(c echo.Context) error {
		return uc.GetAssignTask(c)
	})
	e.POST("/user/lock", func(c echo.Context) error {
		return uc.Lock(c)
	})
	e.POST("/user/change-pass", func(c echo.Context) error {
		return uc.UpdatePass(c)
	})
	e.POST("/user/unlock", func(c echo.Context) error {
		return uc.Unlock(c)
	})
	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Hander room")
		ws.RoomHandler(w, r)
	})
}

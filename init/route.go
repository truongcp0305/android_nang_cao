package init

import (
	"android-service/adapter/controller"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, tc controller.TaskController, uc controller.UserController, wc controller.WordController, ws controller.SocketController) {
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
	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Hander room")
		ws.RoomHandler(w, r)
	})
	//http.HandleFunc("/room/join", ws.RoomHandler)
}

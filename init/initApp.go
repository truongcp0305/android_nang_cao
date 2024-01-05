package init

import (
	"android-service/adapter/connection"
	"android-service/registry"

	"github.com/labstack/echo/v4"
)

func InitApp() {
	mon := connection.Connn()
	r := registry.NewRegistry(mon)
	tc := r.NewTaskController()
	uc := r.NewUserController()
	wc := r.NewWordController()
	sc := r.NewSocketController()
	e := echo.New()
	NewRouter(e, tc, uc, wc, sc)
	//go http.ListenAndServe(":8081", nil)
	go e.Logger.Fatal(e.Start(":8080"))
}

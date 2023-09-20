package init

import (
	"android-service/adapter/connection"
	"android-service/registry"
	"fmt"

	"github.com/labstack/echo/v4"
)

func InitApp() {
	pg := connection.Conn()
	fmt.Println(pg)
	r := registry.NewRegistry(pg)
	tc := r.NewTaskController()
	e := echo.New()
	NewRouter(e, tc)
	e.Logger.Fatal(e.Start(":1323"))
}

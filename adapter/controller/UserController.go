package controller

import (
	"android-service/adapter/incoming"
	"android-service/usecase/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{
		userService: userService,
	}
}

// @Summary Create account
// @Tags User
// @Param userInfo body incoming.CreateUserParam true "userName and password"
// @Success 200 {object} outgoing.UserReturn
// @Failure 400 {object} outgoing.ModelReturn
// @Failure 500 {object} outgoing.ModelReturn
// @Router /user [post]
func (Cs *UserController) Create(c echo.Context) error {
	var params incoming.CreateUserParam
	c.Bind(&params)
	if params.UserName == "" || params.Password == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := params.GetModel()
	userId, err := Cs.userService.CreateAccount(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "create sucsess",
		"userId":  userId,
	})
}

// @Summary Login
// @Tags User
// @Param userInfo body incoming.CreateUserParam true "userName and password"
// @Success 200 {object} outgoing.UserReturn
// @Failure 400 {object} outgoing.ModelReturn
// @Failure 500 {object} outgoing.ModelReturn
// @Router /user/login [post]
func (Cs *UserController) Login(c echo.Context) error {
	var params incoming.LoginParam
	c.Bind(&params)
	if params.UserName == "" || params.Password == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	user := params.GetModel()
	userId, err := Cs.userService.Login(user)
	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login sucsess",
		"userId":  userId,
	})
}

func (Cs *UserController) Logout(c echo.Context) error {
	return nil
}

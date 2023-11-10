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
	result, err := Cs.userService.CreateAccount(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
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
	result, err := Cs.userService.Login(user)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func (Cs *UserController) Update(c echo.Context) error {
	var params incoming.UpdateUserInfoParam
	c.Bind(&params)
	if params.UserId == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	info := params.GetModel()
	err := Cs.userService.Updateinfor(info)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "update sucsess")
}

func (Cs *UserController) Logout(c echo.Context) error {
	return nil
}

func (Cs *UserController) GetList(c echo.Context) error {
	result, err := Cs.userService.GetList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

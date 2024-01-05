package controller

import (
	"android-service/adapter/incoming"
	"android-service/infrastructure"
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
	result.Token, err = infrastructure.CreateJwt(*user)
	if err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
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
	result.Token, err = infrastructure.CreateJwt(*user)
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

func (Cs *UserController) GetAssignTask(c echo.Context) error {
	var param incoming.GetAssignTaskParams
	c.Bind(&param)
	if param.UserId == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	result, err := Cs.userService.GetAssignTasks(param.UserId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": result,
	})
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

func (Cs *UserController) ResetPass(c echo.Context) error {
	params := incoming.ResetPassIncoming{}
	c.Bind(&params)
	if params.Email == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err := Cs.userService.SendMailResetPass(params.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "ok")
}

func (Cs *UserController) Reset(c echo.Context) error {
	var param incoming.ResetLinkIncoming
	c.Bind(&param)
	if param.Value == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err := Cs.userService.ResetPass(param.Value)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Password has been reset!")
}

func (Cs *UserController) UpdatePass(c echo.Context) error {
	var param incoming.UpdatePassParam
	c.Bind(&param)
	if param.NewPass == "" || param.Pass == "" || param.UserName == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err := Cs.userService.UpdatePass(param.GetModel(), param.NewPass)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Password has been Change!")
}

func (Cs *UserController) Lock(c echo.Context) error {
	var param incoming.UpdatePassParam
	c.Bind(&param)
	if param.UserName == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err := Cs.userService.Lock(param.GetModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Lock account succesfuly!")
}

func (Cs *UserController) Unlock(c echo.Context) error {
	var param incoming.UpdatePassParam
	c.Bind(&param)
	if param.UserName == "" {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err := Cs.userService.UnLock(param.GetModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Unlock account succesfuly!")
}

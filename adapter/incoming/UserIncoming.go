package incoming

import (
	"android-service/model"
)

type CreateUserParam struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
}

func (ic *CreateUserParam) GetModel() *model.User {
	return &model.User{
		UserName: ic.UserName,
		Pass:     ic.Password,
	}
}

type LoginParam struct {
	UserName string `json:"userName" form:"userName"`
	Password string `json:"password" form:"password"`
}

func (ic *LoginParam) GetModel() *model.User {
	return &model.User{
		UserName: ic.UserName,
		Pass:     ic.Password,
	}
}

type UpdateUserInfoParam struct {
	UserId     string `json:"userId" form:"userId"`
	UserName   string `json:"userName" form:"userName"`
	Point      string `json:"point" form:"point"`
	OtherInfor string `josn:"otherInfor" form:"otherInfor"`
}

func (u *UpdateUserInfoParam) GetModel() *model.UserInfo {
	return &model.UserInfo{
		UserId:     u.UserId,
		UserName:   u.UserName,
		Point:      u.Point,
		OtherInfor: u.OtherInfor,
	}
}

type StatusIncoming struct {
	Id     string `json:"id" form:"id"`
	Status string `json:"status" form:"status"`
	Point  string `json:"point" form:"point"`
}

func (s *StatusIncoming) GetModel() *model.MatchStatus {
	return &model.MatchStatus{
		Id:     s.Id,
		Status: s.Status,
		Point:  s.Point,
	}
}

type GetAssignTaskParams struct {
	UserId string `json:"userId" form:"userId"`
}

type ResetPassIncoming struct {
	Email string `json:"email" form:"email"`
}

type ResetLinkIncoming struct {
	Value string `param:"encrypt" query:"encrypt"`
}

type UpdatePassParam struct {
	UserName string `json:"userName" form:"userName"`
	Pass     string `json:"password" form:"password"`
	NewPass  string `json:"newPassword" form:"newPassword"`
}

func (u *UpdatePassParam) GetModel() *model.User {
	return &model.User{
		Pass:     u.Pass,
		UserName: u.UserName,
	}
}

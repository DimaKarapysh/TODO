package rest_delivery

import (
	"Portfolio_Nodes/domain"
	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	userUCase domain.UserIter
}

func NewUserDelivery(userIter domain.UserIter) *UserDelivery {
	return &UserDelivery{userUCase: userIter}
}

func (c *UserDelivery) Router(r *gin.RouterGroup) {
	r.POST("/auth/register", c.Register)
	r.POST("/user/login", c.Login)
	r.POST("/user/verify", c.Verify)
}

func (c *UserDelivery) Register(r *gin.Context) {
	form := domain.UserForm{}
	err := r.ShouldBindJSON(&form)
	if err != nil {
		_ = r.Error(NewUserError("Not Invalid Data", err))
		return
	}
	err = c.userUCase.Register(form.Login, form.Pwd)
	if err != nil {
		_ = r.Error(err)
		return
	}
	r.JSON(200, MakeSuccessWithData("Created!"))
}

func (c *UserDelivery) Login(r *gin.Context) {
	form := domain.UserForm{}
	err := r.ShouldBindJSON(&form)
	if err != nil {
		_ = r.Error(NewUserError("Not Invalid Data", err))
		return
	}
	t, err := c.userUCase.Login(form.Login, form.Pwd)
	if err != nil {
		_ = r.Error(err)
		return
	}
	r.JSON(200, t)
}

func (c *UserDelivery) Verify(r *gin.Context) {
	token := domain.Token{}
	err := r.ShouldBindJSON(&token)
	if err != nil {
		_ = r.Error(NewUserError("Not Invalid Data", err))
	}
	user, err := c.userUCase.Verify(token.Token)
	if err != nil {
		_ = r.Error(err)
		return
	}
	r.JSON(200, user)
}

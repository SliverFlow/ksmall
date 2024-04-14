package service

import (
	"github.com/SliverFlow/ksmall/app/user/http/internal/biz"
	"github.com/SliverFlow/ksmall/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserService struct {
	logger      *zap.Logger
	userUsecase *biz.UserUsecase
}

func NewUserService(logger *zap.Logger, userUsecase *biz.UserUsecase) *UserService {
	return &UserService{logger: logger, userUsecase: userUsecase}
}

func (us *UserService) InitUserRouter(r *gin.RouterGroup) {
	r.GET("find", us.FindById)
}

func (us *UserService) FindById(c *gin.Context) {

	userInfoReply, err := us.userUsecase.FindById(c, 9)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(userInfoReply, c)
}

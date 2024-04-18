package service

import (
	"github.com/SliverFlow/ksmall/app/user/http/internal/biz"
	"github.com/SliverFlow/ksmall/app/user/http/internal/model/request"
	"github.com/SliverFlow/ksmall/common"
	"github.com/SliverFlow/ksmall/common/response"
	"github.com/SliverFlow/ksmall/common/util"
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
	r.GET("find", us.Find)        // 根据id查询用户
	r.POST("login", us.Login)     // 登录
	r.POST("captcha", us.Captcha) // 获取验证码
}

// Find
// @Author: [github.com/SliverFlow]
// @Desc: 根据id查询用户
func (us *UserService) Find(c *gin.Context) {
	var req common.IdReq
	if err := c.ShouldBind(&req); err != nil {
		us.logger.Error("[http] user find by id params error", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}
	us.logger.Info("[http] user login request", zap.Any("req", req))

	userInfoReply, err := us.userUsecase.Find(c, req.Id)
	if err != nil {
		us.logger.Error("[http] user find by id error", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(userInfoReply, c)
}

// Login
// @Author: [github.com/SliverFlow]
// @Desc: 登录
func (us *UserService) Login(c *gin.Context) {
	var req request.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		us.logger.Error("[http] user login params error", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	us.logger.Info("[http] user login request", zap.Any("req", req))
	userInfo, err := us.userUsecase.Login(c, &req)
	if err != nil {
		us.logger.Error("[http] user login error", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(userInfo, c)
}

// Captcha
// @Author: [github.com/SliverFlow]
// @Desc: 获取验证码
func (us *UserService) Captcha(c *gin.Context) {
	var req request.CaptchaReq
	if err := c.ShouldBindJSON(&req); err != nil {
		us.logger.Error("[http] user captcha params error", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	us.logger.Info("[http] user captcha request", zap.Any("req", req))

	if err := us.userUsecase.Captcha(c, req.Account, req.Type); err != nil {
		us.logger.Error("[http] user captcha error", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

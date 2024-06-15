package service

import (
	"github.com/SliverFlow/ksmall/monserver/internal/biz"
	"github.com/SliverFlow/ksmall/monserver/tracing"
	"go.uber.org/zap"
)

type AuthorityService struct {
	tracing.Service
	logger  *zap.Logger
	usecase *biz.AuthorityUsecase
}

func NewAuthorityService(usecase *biz.AuthorityUsecase, logger *zap.Logger) *AuthorityService {
	return &AuthorityService{
		usecase: usecase,
		logger:  logger,
	}
}

package service

import (
	"github.com/SliverFlow/ksmall/monserver/internal/biz"
	"github.com/SliverFlow/ksmall/monserver/tracing"
	"go.uber.org/zap"
)

type AuthorityGroupService struct {
	tracing.Service
	logger  *zap.Logger
	usecase *biz.AuthorityGroupUsecase
}

func NewAuthorityGroupService(usecase *biz.AuthorityGroupUsecase, logger *zap.Logger) *AuthorityGroupService {
	return &AuthorityGroupService{
		usecase: usecase,
		logger:  logger,
	}
}

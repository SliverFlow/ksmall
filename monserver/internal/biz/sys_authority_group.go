package biz

import (
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/tracing"
	"go.uber.org/zap"
)

type AuthorityGroupUsecase struct {
	tracing.Biz
	logger             *zap.Logger
	authorityGroupRepo repo.AuthorityGroupRepo
}

func NewAuthorityGroupUsecase(authorityGroupRepo repo.AuthorityGroupRepo) *AuthorityGroupUsecase {
	return &AuthorityGroupUsecase{
		authorityGroupRepo: authorityGroupRepo,
	}
}

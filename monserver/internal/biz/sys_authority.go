package biz

import (
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/tracing"
	"go.uber.org/zap"
)

type AuthorityUsecase struct {
	tracing.Biz
	logger        *zap.Logger
	authorityRepo repo.AuthorityRepo
}

func NewAuthorityUsecase(authorityRepo repo.AuthorityRepo) *AuthorityUsecase {
	return &AuthorityUsecase{
		authorityRepo: authorityRepo,
	}
}

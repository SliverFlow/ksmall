package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/data/repo"
)

type UserUsecase struct {
	iUserRepo repo.IUserRepo
}

func NewUserUsecase(iUserRepo repo.IUserRepo) *UserUsecase {
	return &UserUsecase{iUserRepo: iUserRepo}
}

func (uu *UserUsecase) FindByUsername(ctx context.Context, username string) {

}

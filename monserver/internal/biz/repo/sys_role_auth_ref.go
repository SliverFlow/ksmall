package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
)

type RoleAuthRefRepo interface {
	FindByRoleId(ctx context.Context, roleId int64) ([]*model.RoleAuthRef, error)
}

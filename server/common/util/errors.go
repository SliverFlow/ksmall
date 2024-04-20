package util

import (
	"github.com/SliverFlow/ksmall/common/constant"
	"google.golang.org/grpc/status"
)

func IsGormNotFoundErr(err error) bool {
	if fromError, ok := status.FromError(err); ok {
		if code := uint32(fromError.Code()); code == constant.GormNotFoundCode {
			return true
		}
	}
	return false
}

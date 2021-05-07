package errorutil

import (
	"github.com/dajinkuang/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func IsError(err error) bool {
	err = errors.Cause(err)
	if s, ok := status.FromError(err); ok {
		if s.Code() > 100 {
			return false
		}
		switch s.Code() {
		case codes.InvalidArgument, codes.Unauthenticated, codes.FailedPrecondition:
			return false
		default:
			return true
		}
	}
	if err != nil {
		return true
	}
	return false
}

package errorutil

import (
	"database/sql"

	"github.com/dajinkuang/errors"
	"github.com/gomodule/redigo/redis"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ChangeErr2Grpc 将 error 转为 grpc error
func ChangeErr2Grpc(err error) error {
	if err == nil {
		return nil
	}
	err = errors.Cause(err)
	if _, ok := status.FromError(err); ok {
		return err
	}
	if err == sql.ErrNoRows {
		return status.Errorf(codes.NotFound, "Sorry, Resources not found")
	}
	if err == redis.ErrNil {
		return status.Errorf(codes.NotFound, "Sorry ~ Resources not found")
	}
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	return err
}

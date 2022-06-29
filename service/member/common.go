package member

import (
	"context"
)

var (
	ctx    = context.TODO()
	ctxb   = context.Background()
	jwtKey = []byte("Secret")
)

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

package post

import (
	"context"
)

var jwtKey = []byte("Secret")

var ctxb = context.Background()

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

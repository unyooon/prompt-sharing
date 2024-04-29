package main

import (
	"github.com/google/wire"
	"github.com/unyooon/prompt-sharing/domain/setting"
)

func InitializeHandler(s setting.Setting) *Routing {
	wire.Build(SuperSet, NewRouting)
	return nil
}

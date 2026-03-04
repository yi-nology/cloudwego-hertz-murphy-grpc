package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	common "{{.module_name}}/gen/http/router/common"
)

func GeneratedRegister(r *server.Hertz) {
	common.Register(r)
}

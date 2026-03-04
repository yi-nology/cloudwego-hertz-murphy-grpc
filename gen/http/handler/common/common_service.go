package common

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"{{.module_name}}/gen/http/model/common"
	"{{.module_name}}/internal/conf"
	"{{.module_name}}/internal/pkg/resp"
)

func Health(ctx context.Context, c *app.RequestContext) {
	var req common.EmptyReq
	if err := c.BindAndValidate(&req); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	resp.Success(c, &common.HealthResp{Status: "ok"})
}

func Index(ctx context.Context, c *app.RequestContext) {
	var req common.EmptyReq
	if err := c.BindAndValidate(&req); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	cfg := conf.GlobalConfig
	resp.Success(c, &common.IndexResp{
		Name:    cfg.App.Name,
		Version: cfg.App.Version,
		Status:  "running",
	})
}

func Ping(ctx context.Context, c *app.RequestContext) {
	var req common.EmptyReq
	if err := c.BindAndValidate(&req); err != nil {
		resp.BadRequest(c, err.Error())
		return
	}

	resp.Success(c, &common.PingResp{Message: "pong"})
}

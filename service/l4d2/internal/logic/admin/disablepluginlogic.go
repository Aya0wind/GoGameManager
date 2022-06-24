package admin

import (
	"context"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisablePluginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDisablePluginLogic(ctx context.Context, svcCtx *svc.ServiceContext) DisablePluginLogic {
	return DisablePluginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DisablePluginLogic) DisablePlugin(req types.DisablePluginRequest) (resp *types.DisablePluginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

package admin

import (
	"context"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnablePluginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnablePluginLogic(ctx context.Context, svcCtx *svc.ServiceContext) EnablePluginLogic {
	return EnablePluginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnablePluginLogic) EnablePlugin(req types.EnablePluginRequest) (resp *types.EnablePluginResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

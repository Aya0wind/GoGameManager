package admin

import (
	"context"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerPluginsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServerPluginsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetServerPluginsLogic {
	return GetServerPluginsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServerPluginsLogic) GetServerPlugins(req types.GetServerPluginsRequest) (resp *types.GetServerPluginsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}

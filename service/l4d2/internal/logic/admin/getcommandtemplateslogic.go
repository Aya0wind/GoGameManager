package admin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"
)

type GetCommandTemplatesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommandTemplatesLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCommandTemplatesLogic {
	return GetCommandTemplatesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommandTemplatesLogic) GetCommandTemplates(req types.GetCommandTemplateReqeust) (resp *types.GetCommandTemplateResponse, err error) {
	resp = &types.GetCommandTemplateResponse{
		Code: 200,
		Msg:  "ok",
		Data: l.svcCtx.Templates,
	}
	return
}

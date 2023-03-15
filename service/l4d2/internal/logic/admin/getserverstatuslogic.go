package admin

import (
	"context"
	"l4d2/common"
	"l4d2/service/l4d2/internal/logic/admin/utils"
	"net/http"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetServerStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetServerStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetServerStatusLogic {
	return GetServerStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetServerStatusLogic) GetServerStatus(_ types.GetServerStatusRequest) (resp *types.GetServerStatusResponse, err error) {
	result, err := common.ExecRconCommand(&l.svcCtx.Config.Rcon, "status")
	if err != nil {
		l.Logger.Error(err)
		resp = &types.GetServerStatusResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}
		err = nil
		return
	}
	info := utils.ParseServerInfo(result)
	resp = &types.GetServerStatusResponse{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: info,
	}
	return
}

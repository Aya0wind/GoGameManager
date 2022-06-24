package admin

import (
	"context"
	"l4d2/service/l4d2/internal/logic/admin/impl"
	"net/http"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExecuteCommandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExecuteCommandLogic(ctx context.Context, svcCtx *svc.ServiceContext) ExecuteCommandLogic {
	return ExecuteCommandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExecuteCommandLogic) ExecuteCommand(req types.ExecuteCommandRequest) (resp *types.ExecuteCommandResponse, err error) {
	result, err := impl.ExecRconCommand(&l.svcCtx.Config.Rcon, req.Command)
	l.Logger.Info(result)
	if err != nil {
		resp = &types.ExecuteCommandResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}
		err = nil
		return
	}
	resp = &types.ExecuteCommandResponse{
		Code: http.StatusOK,
		Msg:  "ok",
	}
	return
}

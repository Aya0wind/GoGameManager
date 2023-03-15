package admin

import (
	"context"
	"net/http"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMapGroupByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMapGroupByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMapGroupByIDLogic {
	return GetMapGroupByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMapGroupByIDLogic) GetMapGroupByID(req types.GetMapGroupByIDRequest) (resp *types.GetMapGroupByIDResponse, err error) {
	mapGroup, err := l.svcCtx.Db.QueryMapGroupByID(req.ID)
	if err != nil {
		l.Logger.Error(err)
		resp = &types.GetMapGroupByIDResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}
		err = nil
		return
	}
	var Data types.MapGroup
	Data.Id = mapGroup.Id
	if mapGroup.Name.Valid {
		Data.Name = mapGroup.Name.String
	} else {
		Data.Name = "无法解析到地图名称"
	}

	if mapGroup.StartName.Valid {
		Data.StartName = mapGroup.StartName.String
	} else {
		Data.StartName = "无法解析到关卡名称"
	}
	resp = &types.GetMapGroupByIDResponse{
		Code: http.StatusBadRequest,
		Msg:  "ok",
		Data: &Data,
	}
	return
}

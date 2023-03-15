package admin

import (
	"context"
	"net/http"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMapGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMapGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMapGroupLogic {
	return GetMapGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMapGroupLogic) GetMapGroup(_ types.GetMapGroupRequest) (resp *types.GetMapGroupResponse, err error) {
	mapGroups, err := l.svcCtx.Db.QueryAllMapGroup()
	if err != nil {
		l.Logger.Error(err)
		resp = &types.GetMapGroupResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}
		err = nil
		return
	}
	MapGroups := make([]types.MapGroup, 0)
	for _, mapGroup := range mapGroups {
		var data types.MapGroup
		data.Id = mapGroup.Id
		data.UpdatedAt = mapGroup.UpdatedTime
		data.CreatedAt = mapGroup.CreatedTime
		data.LastPlayTime = mapGroup.LastPlayTime
		if mapGroup.Name.Valid {
			data.Name = mapGroup.Name.String
		} else {
			data.Name = "无法解析到地图名称"
		}

		if mapGroup.StartName.Valid {
			data.StartName = mapGroup.StartName.String
		} else {
			data.StartName = "无法解析到关卡名称"
		}
		MapGroups = append(MapGroups, data)
	}

	resp = &types.GetMapGroupResponse{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: MapGroups,
	}
	return
}

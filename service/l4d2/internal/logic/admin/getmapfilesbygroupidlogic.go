package admin

import (
	"context"
	"net/http"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMapFilesByGroupIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMapFilesByGroupIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetMapFilesByGroupIDLogic {
	return GetMapFilesByGroupIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMapFilesByGroupIDLogic) GetMapFilesByGroupID(req types.GetMapFilesByGroupIDRequest) (resp *types.GetMapFilesByGroupIDResponse, err error) {
	mapFiles, err := l.svcCtx.Db.QueryMapFileByMapGroupID(req.ID)
	if err != nil {
		resp = &types.GetMapFilesByGroupIDResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}
		err = nil
		return
	}
	mapFilesResp := make([]types.MapFile, 0)
	for _, mapFile := range mapFiles {
		mapFilesResp = append(mapFilesResp, types.MapFile{
			Id:        mapFile.Id,
			FileName:  mapFile.FileName,
			GroupID:   mapFile.GroupID,
			CreatedAt: mapFile.CreatedAt,
			UpdatedAt: mapFile.UpdatedAt,
		})
	}
	resp = &types.GetMapFilesByGroupIDResponse{
		Code: http.StatusOK,
		Msg:  "ok",
		Data: mapFilesResp,
	}
	return
}

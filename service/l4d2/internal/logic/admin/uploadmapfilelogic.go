package admin

import (
	"context"
	"database/sql"
	"fmt"
	"l4d2/service/l4d2/internal/logic/admin/utils"
	"l4d2/service/l4d2/internal/logic/admin/utils/vdfparser"
	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"
	"l4d2/service/l4d2/model"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadMapFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadMapFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadMapFileLogic {
	return UploadMapFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadMapFileLogic) UploadMapFile(req types.UploadMapFileRequest, fileHeaders []*multipart.FileHeader) (resp *types.UploadMapFileResponse, err error) {
	groupID := req.GroupID
	for _, file := range fileHeaders {
		path := fmt.Sprintf("%s/%s", l.svcCtx.Config.Path.VpkPath, file.Filename)
		if utils.IsExist(path) {
			continue
		}
		err = utils.SaveMultipartFile(file, path)
		if err != nil {
			resp = &types.UploadMapFileResponse{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
			}
			err = nil
			return
		}
		mapFileInfo, err := vdfparser.ParseVpkMapFileInfo(path)
		var mapGroup *model.MapGroup
		if err != nil {
			mapGroup = &model.MapGroup{
				Id:           groupID,
				Name:         sql.NullString{},
				StartName:    sql.NullString{},
				LastPlayTime: time.Now().Unix(),
				CreatedTime:  time.Now().Unix(),
				UpdatedTime:  time.Now().Unix(),
			}
		} else {
			mapGroup = &model.MapGroup{
				Id: groupID,
				Name: sql.NullString{
					String: mapFileInfo[0].Name,
					Valid:  true,
				},
				StartName: sql.NullString{
					String: mapFileInfo[0].StartName,
					Valid:  true,
				},
				LastPlayTime: time.Now().Unix(),
				CreatedTime:  time.Now().Unix(),
				UpdatedTime:  time.Now().Unix(),
			}
		}
		err = l.svcCtx.Db.InsertMapGroupAndMapFile(mapGroup, file.Filename)
		if err != nil {
			resp = &types.UploadMapFileResponse{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
			}
			err = nil
			return resp, err
		}
	}
	resp = &types.UploadMapFileResponse{
		Code: http.StatusOK,
		Msg:  "ok",
	}
	return
}

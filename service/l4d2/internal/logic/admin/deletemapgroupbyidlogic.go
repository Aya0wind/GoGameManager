package admin

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMapGroupByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMapGroupByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteMapGroupByIDLogic {
	return DeleteMapGroupByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMapGroupByIDLogic) DeleteMapGroupByID(req types.DeleteMapGroupByIDRequest) (resp *types.DeleteMapGroupByIDResponse, err error) {
	needDeleteFiles, err := l.svcCtx.Db.DeleteMapGroupAndFilesByID(req.ID)
	//删除所有地图文件
	for _, file := range needDeleteFiles {
		err = os.Remove(fmt.Sprintf("%s/%s", l.svcCtx.Config.Path.VpkPath, file.FileName))
		if err != nil {
			l.Logger.Error(err)
		}
	}
	if err != nil {
		resp = &types.DeleteMapGroupByIDResponse{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		}
		err = nil
		return
	}

	resp = &types.DeleteMapGroupByIDResponse{
		Code: http.StatusOK,
		Msg:  "ok",
	}
	return
}

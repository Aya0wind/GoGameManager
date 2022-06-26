package admin

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"l4d2/service/l4d2/internal/logic/admin"
	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"
	"net/http"
)

func UploadMapFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadMapFileRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		fileHeader := r.MultipartForm.File["file"]
		l := admin.NewUploadMapFileLogic(r.Context(), svcCtx)
		resp, err := l.UploadMapFile(req, fileHeader)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package admin

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"l4d2/service/l4d2/internal/logic/admin"
	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"
)

func UploadMapFileNoVerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadMapFileNoVerifyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		fileHeader := r.MultipartForm.File["file"]
		groupIDHeader := r.Header["Group-ID"]
		var groupID int64
		if len(groupIDHeader) == 0 {
			groupID = 0
		} else {
			number, err := strconv.ParseInt(groupIDHeader[0], 10, 64)
			groupID = number
			if err != nil {
				httpx.Error(w, err)
				return
			}
		}
		l := admin.NewUploadMapFileNoVerifyLogic(r.Context(), svcCtx)
		resp, err := l.UploadMapFileNoVerify(fileHeader, groupID)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

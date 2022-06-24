package admin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"l4d2/service/l4d2/internal/logic/admin"
	"l4d2/service/l4d2/internal/svc"
	"l4d2/service/l4d2/internal/types"
)

func DisablePluginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DisablePluginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := admin.NewDisablePluginLogic(r.Context(), svcCtx)
		resp, err := l.DisablePlugin(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

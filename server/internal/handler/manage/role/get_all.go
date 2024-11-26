package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"server/internal/logic/manage/role"
	"server/internal/svc"
	types "server/internal/types/manage/role"
)

func GetAll(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewGetAll(r.Context(), svcCtx)
		resp, err := l.GetAll(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

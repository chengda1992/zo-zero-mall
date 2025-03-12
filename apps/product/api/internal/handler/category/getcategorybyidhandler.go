package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_mall/apps/product/api/internal/logic/category"
	"go_mall/apps/product/api/internal/svc"
	"go_mall/apps/product/api/internal/types"
)

// 分类详情
func GetCategoryByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCategoryByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewGetCategoryByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryById(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

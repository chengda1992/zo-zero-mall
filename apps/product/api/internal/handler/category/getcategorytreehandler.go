package category

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_mall/apps/product/api/internal/logic/category"
	"go_mall/apps/product/api/internal/svc"
)

// 分类树
func GetCategoryTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewGetCategoryTreeLogic(r.Context(), svcCtx)
		resp, err := l.GetCategoryTree()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

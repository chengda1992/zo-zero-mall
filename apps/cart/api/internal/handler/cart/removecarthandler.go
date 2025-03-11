package cart

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_mall/apps/cart/api/internal/logic/cart"
	"go_mall/apps/cart/api/internal/svc"
	"go_mall/apps/cart/api/internal/types"
)

// 删除购物车
func RemoveCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cart.NewRemoveCartLogic(r.Context(), svcCtx)
		resp, err := l.RemoveCart(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

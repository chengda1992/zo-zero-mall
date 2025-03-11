// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package handler

import (
	"net/http"

	cart "go_mall/apps/cart/api/internal/handler/cart"
	"go_mall/apps/cart/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// 获取购物车列表
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: cart.ListCartHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/cart/cart"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 添加购物车
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: cart.AddCartHandler(serverCtx),
			},
			{
				// 删除购物车
				Method:  http.MethodPost,
				Path:    "/remove",
				Handler: cart.RemoveCartHandler(serverCtx),
			},
			{
				// 更新购物车
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: cart.UpdateCartHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/v1/cart/cart"),
	)
}

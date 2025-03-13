// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package types

type AddCategoryReq struct {
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
}

type AddProductReq struct {
	CateId   int64    `json:"cate_id"`
	Name     string   `json:"name"`
	Detail   string   `json:"detail"`
	Subtitle string   `json:"subtitle"`
	Price    float64  `json:"price"`
	Stock    int64    `json:"stock"`
	Status   int64    `json:"status"`
	Images   []string `json:"images"`
}

type AddProductResp struct {
	Id int64 `json:"id"`
}

type CategoryNode struct {
	Id       int64          `json:"id"`
	Name     string         `json:"name"`
	Children []CategoryNode `json:"children"`
}

type DeleteCategoryReq struct {
	Id int64 `json:"id"`
}

type DeleteProductReq struct {
	Id int64 `json:"id"`
}

type GetCategoryByIdReq struct {
	Id int64 `json:"id"`
}

type GetProductReq struct {
	Id int64 `json:"id"`
}

type Response struct {
	Code    int64       `json:"code"`    // 状态码
	Message string      `json:"message"` // 返回信息
	Data    interface{} `json:"data"`    // 返回数据
}

type SearchProductReq struct {
	Keyword string `json:"keyword"`
	Page    int64  `json:"page"`
	Size    int64  `json:"size"`
}

type UpdateCategoryReq struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

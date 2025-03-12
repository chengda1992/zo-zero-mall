package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CategoryModel = (*customCategoryModel)(nil)

type (
	// CategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoryModel.
	CategoryModel interface {
		categoryModel
		FindOneByNameAndParentId(ctx context.Context, name string, parentId int64) (*Category, error)
	}

	customCategoryModel struct {
		*defaultCategoryModel
	}
)

// NewCategoryModel returns a model for the database table.
func NewCategoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CategoryModel {
	return &customCategoryModel{
		defaultCategoryModel: newCategoryModel(conn, c, opts...),
	}
}

// 根据名称和父级ID查询分类
func (m *defaultCategoryModel) FindOneByNameAndParentId(ctx context.Context, name string, parentId int64) (*Category, error) {
	sql := "select id, parentid,name,status,create_time,update_time from category where name = ? and parentid = ?"
	var resp Category
	err := m.QueryRowNoCacheCtx(ctx, &resp, sql, name, parentId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

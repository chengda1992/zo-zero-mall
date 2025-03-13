// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.8.1

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	productFieldNames          = builder.RawFieldNames(&Product{})
	productRows                = strings.Join(productFieldNames, ",")
	productRowsExpectAutoSet   = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	productRowsWithPlaceHolder = strings.Join(stringx.Remove(productFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	productModel interface {
		Insert(ctx context.Context, data *Product) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Product, error)
		Update(ctx context.Context, data *Product) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultProductModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Product struct {
		Id         uint64         `db:"id"`          // 商品id
		Cateid     uint64         `db:"cateid"`      // 类别Id
		Name       string         `db:"name"`        // 商品名称
		Subtitle   string         `db:"subtitle"`    // 商品副标题
		Images     sql.NullString `db:"images"`      // 图片地址,json格式,扩展用
		Detail     sql.NullString `db:"detail"`      // 商品详情
		Price      float64        `db:"price"`       // 价格,单位-元保留两位小数
		Stock      int64          `db:"stock"`       // 库存数量
		Status     int64          `db:"status"`      // 商品状态.1-在售 2-下架 3-删除
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime time.Time      `db:"update_time"` // 更新时间
	}
)

func newProductModel(conn sqlx.SqlConn) *defaultProductModel {
	return &defaultProductModel{
		conn:  conn,
		table: "`product`",
	}
}

func (m *defaultProductModel) Delete(ctx context.Context, id uint64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultProductModel) FindOne(ctx context.Context, id uint64) (*Product, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", productRows, m.table)
	var resp Product
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultProductModel) Insert(ctx context.Context, data *Product) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, productRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Cateid, data.Name, data.Subtitle, data.Images, data.Detail, data.Price, data.Stock, data.Status)
	return ret, err
}

func (m *defaultProductModel) Update(ctx context.Context, data *Product) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, productRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Cateid, data.Name, data.Subtitle, data.Images, data.Detail, data.Price, data.Stock, data.Status, data.Id)
	return err
}

func (m *defaultProductModel) tableName() string {
	return m.table
}

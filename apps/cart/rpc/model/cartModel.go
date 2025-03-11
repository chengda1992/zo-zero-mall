package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CartModel = (*customCartModel)(nil)

type (
	// CartModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCartModel.
	CartModel interface {
		cartModel
		withSession(session sqlx.Session) CartModel
		Transact(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
	}

	customCartModel struct {
		*defaultCartModel
	}
)

// NewCartModel returns a model for the database table.
func NewCartModel(conn sqlx.SqlConn) CartModel {
	return &customCartModel{
		defaultCartModel: newCartModel(conn),
	}
}

func (m *customCartModel) withSession(session sqlx.Session) CartModel {
	return NewCartModel(sqlx.NewSqlConnFromSession(session))
}

// Transact 事务支持
func (m *customCartModel) Transact(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

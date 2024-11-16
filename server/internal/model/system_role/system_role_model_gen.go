// Code generated by goctl. DO NOT EDIT.

package system_role

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	systemRoleFieldNames          = builder.RawFieldNames(&SystemRole{})
	systemRoleRows                = strings.Join(systemRoleFieldNames, ",")
	systemRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(systemRoleFieldNames, "`id`"), ",")
	systemRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(systemRoleFieldNames, "`id`"), "=?,") + "=?"
)

type (
	systemRoleModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *SystemRole) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id uint64) (*SystemRole, error)
		Update(ctx context.Context, session sqlx.Session, data *SystemRole) error
		Delete(ctx context.Context, session sqlx.Session, id uint64) error

		// custom interface generated by jzero
		BulkInsert(ctx context.Context, session sqlx.Session, datas []*SystemRole) error
		FindByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*SystemRole, error)
		FindOneByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) (*SystemRole, error)
		PageByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*SystemRole, int64, error)
		UpdateFieldsByCondition(ctx context.Context, session sqlx.Session, field map[string]any, conds ...condition.Condition) error
		DeleteByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) error
	}

	defaultSystemRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SystemRole struct {
		Id         uint64        `db:"id"`
		CreateTime time.Time     `db:"create_time"`
		UpdateTime time.Time     `db:"update_time"`
		CreateBy   sql.NullInt64 `db:"create_by"`
		UpdateBy   sql.NullInt64 `db:"update_by"`
		Name       string        `db:"name"`
		Status     string        `db:"status"`
		Code       string        `db:"code"`
		Desc       string        `db:"desc"`
	}
)

func newSystemRoleModel(conn sqlx.SqlConn) *defaultSystemRoleModel {
	return &defaultSystemRoleModel{
		conn:  conn,
		table: "`system_role`",
	}
}

func (m *defaultSystemRoleModel) Delete(ctx context.Context, session sqlx.Session, id uint64) error {
	sb := sqlbuilder.DeleteFrom(m.table)
	sb.Where(sb.EQ("`id`", id))
	sql, args := sb.Build()
	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, sql, args...)
	} else {
		_, err = m.conn.ExecCtx(ctx, sql, args...)
	}

	return err
}

func (m *defaultSystemRoleModel) FindOne(ctx context.Context, session sqlx.Session, id uint64) (*SystemRole, error) {
	sb := sqlbuilder.Select(systemRoleRows).From(m.table)
	sb.Where(sb.EQ("`id`", id))
	sb.Limit(1)
	sql, args := sb.Build()
	var resp SystemRole
	var err error
	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, sql, args...)
	} else {
		err = m.conn.QueryRowCtx(ctx, &resp, sql, args...)
	}
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSystemRoleModel) Insert(ctx context.Context, session sqlx.Session, data *SystemRole) (sql.Result, error) {
	statement, args := sqlbuilder.NewInsertBuilder().
		InsertInto(m.table).
		Cols(systemRoleRowsExpectAutoSet).
		Values(data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.Name, data.Status, data.Code, data.Desc).Build()
	if session != nil {
		return session.ExecCtx(ctx, statement, args...)
	}
	return m.conn.ExecCtx(ctx, statement, args...)
}

func (m *defaultSystemRoleModel) Update(ctx context.Context, session sqlx.Session, data *SystemRole) error {
	sb := sqlbuilder.Update(m.table)
	split := strings.Split(systemRoleRowsExpectAutoSet, ",")
	var assigns []string
	for _, s := range split {
		assigns = append(assigns, sb.Assign(s, nil))
	}
	sb.Set(assigns...)
	sb.Where(sb.EQ("`id`", nil))
	statement, _ := sb.Build()

	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, statement, data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.Name, data.Status, data.Code, data.Desc, data.Id)
	} else {
		_, err = m.conn.ExecCtx(ctx, statement, data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.Name, data.Status, data.Code, data.Desc, data.Id)
	}
	return err
}

func (m *defaultSystemRoleModel) tableName() string {
	return m.table
}

func (m *customSystemRoleModel) BulkInsert(ctx context.Context, session sqlx.Session, datas []*SystemRole) error {
	sb := sqlbuilder.InsertInto(m.table)
	sb.Cols(systemRoleRowsExpectAutoSet)
	for _, data := range datas {
		sb.Values(data.CreateTime, data.UpdateTime, data.CreateBy, data.UpdateBy, data.Name, data.Status, data.Code, data.Desc)
	}
	statement, args := sb.Build()

	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, statement, args...)
	} else {
		_, err = m.conn.ExecCtx(ctx, statement, args...)
	}
	return err
}

func (m *customSystemRoleModel) FindByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*SystemRole, error) {
	sb := sqlbuilder.Select(systemRoleFieldNames...).From(m.table)
	condition.ApplySelect(sb, conds...)
	statement, args := sb.Build()

	var resp []*SystemRole
	var err error

	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, statement, args...)
	} else {
		err = m.conn.QueryRowsCtx(ctx, &resp, statement, args...)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *customSystemRoleModel) FindOneByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) (*SystemRole, error) {
	sb := sqlbuilder.Select(systemRoleFieldNames...).From(m.table)

	condition.ApplySelect(sb, conds...)
	sb.Limit(1)
	statement, args := sb.Build()

	var resp SystemRole
	var err error

	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, statement, args...)
	} else {
		err = m.conn.QueryRowCtx(ctx, &resp, statement, args...)
	}
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *customSystemRoleModel) PageByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) ([]*SystemRole, int64, error) {
	sb := sqlbuilder.Select(systemRoleFieldNames...).From(m.table)
	countsb := sqlbuilder.Select("count(*)").From(m.table)

	condition.ApplySelect(sb, conds...)

	var countConds []condition.Condition
	for _, cond := range conds {
		if cond.Operator != condition.Limit && cond.Operator != condition.Offset {
			countConds = append(countConds, cond)
		}
	}
	condition.ApplySelect(countsb, countConds...)

	var resp []*SystemRole
	var err error

	statement, args := sb.Build()

	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, statement, args...)
	} else {
		err = m.conn.QueryRowsCtx(ctx, &resp, statement, args...)
	}
	if err != nil {
		return nil, 0, err
	}

	var total int64
	statement, args = countsb.Build()
	if session != nil {
		err = session.QueryRowCtx(ctx, &total, statement, args...)
	} else {
		err = m.conn.QueryRowCtx(ctx, &total, statement, args...)
	}
	if err != nil {
		return nil, 0, err
	}

	return resp, total, nil
}

func (m *customSystemRoleModel) UpdateFieldsByCondition(ctx context.Context, session sqlx.Session, field map[string]any, conds ...condition.Condition) error {
	if field == nil {
		return nil
	}

	sb := sqlbuilder.Update(m.table)
	condition.ApplyUpdate(sb, conds...)

	var assigns []string
	for key, value := range field {
		assigns = append(assigns, sb.Assign(key, value))
	}
	sb.Set(assigns...)

	statement, args := sb.Build()

	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, statement, args...)
	} else {
		_, err = m.conn.ExecCtx(ctx, statement, args...)
	}
	if err != nil {
		return err
	}
	return nil
}

func (m *customSystemRoleModel) DeleteByCondition(ctx context.Context, session sqlx.Session, conds ...condition.Condition) error {
	if len(conds) == 0 {
		return nil
	}
	sb := sqlbuilder.DeleteFrom(m.table)
	condition.ApplyDelete(sb, conds...)
	statement, args := sb.Build()

	var err error
	if session != nil {
		_, err = session.ExecCtx(ctx, statement, args...)
	} else {
		_, err = m.conn.ExecCtx(ctx, statement, args...)
	}
	return err
}

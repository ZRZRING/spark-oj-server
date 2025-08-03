// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContestDao is the data access object for the table contest.
type ContestDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ContestColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ContestColumns defines and stores column names for the table contest.
type ContestColumns struct {
	Cid         string //
	Title       string //
	Password    string //
	Description string //
	StartTime   string //
	EndTime     string //
	LockTime    string //
	CreateBy    string //
	CreateAt    string //
	UpdateAt    string //
	DeleteAt    string //
	Practice    string //
}

// contestColumns holds the columns for the table contest.
var contestColumns = ContestColumns{
	Cid:         "cid",
	Title:       "title",
	Password:    "password",
	Description: "description",
	StartTime:   "start_time",
	EndTime:     "end_time",
	LockTime:    "lock_time",
	CreateBy:    "create_by",
	CreateAt:    "create_at",
	UpdateAt:    "update_at",
	DeleteAt:    "delete_at",
	Practice:    "practice",
}

// NewContestDao creates and returns a new DAO object for table data access.
func NewContestDao(handlers ...gdb.ModelHandler) *ContestDao {
	return &ContestDao{
		group:    "default",
		table:    "contest",
		columns:  contestColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ContestDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ContestDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ContestDao) Columns() ContestColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ContestDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ContestDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ContestDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

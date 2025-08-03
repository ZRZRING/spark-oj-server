// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SubmissionDao is the data access object for the table submission.
type SubmissionDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SubmissionColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SubmissionColumns defines and stores column names for the table submission.
type SubmissionColumns struct {
	Sid        string //
	Title      string //
	Pid        string //
	Username   string //
	Cid        string //
	Result     string //
	Language   string //
	MemoryCost string //
	TimeCost   string //
	Code       string //
	CreateAt   string //
	UpdateAt   string //
	DeleteAt   string //
}

// submissionColumns holds the columns for the table submission.
var submissionColumns = SubmissionColumns{
	Sid:        "sid",
	Title:      "title",
	Pid:        "pid",
	Username:   "username",
	Cid:        "cid",
	Result:     "result",
	Language:   "language",
	MemoryCost: "memory_cost",
	TimeCost:   "time_cost",
	Code:       "code",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeleteAt:   "delete_at",
}

// NewSubmissionDao creates and returns a new DAO object for table data access.
func NewSubmissionDao(handlers ...gdb.ModelHandler) *SubmissionDao {
	return &SubmissionDao{
		group:    "default",
		table:    "submission",
		columns:  submissionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SubmissionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SubmissionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SubmissionDao) Columns() SubmissionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SubmissionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SubmissionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SubmissionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

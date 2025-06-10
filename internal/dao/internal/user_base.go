// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserBaseDao is the data access object for the table user_base.
type UserBaseDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserBaseColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserBaseColumns defines and stores column names for the table user_base.
type UserBaseColumns struct {
	Username   string //
	Password   string //
	Nickname   string //
	Role       string //
	Solved     string //
	Submitted  string //
	Rating     string //
	CfId       string //
	AtcId      string //
	Company    string //
	Department string //
	Major      string //
	Class      string //
	Email      string //
	Tel        string //
	Avatar     string //
	CreateAt   string //
	UpdateAt   string //
	DeleteAt   string //
}

// userBaseColumns holds the columns for the table user_base.
var userBaseColumns = UserBaseColumns{
	Username:   "username",
	Password:   "password",
	Nickname:   "nickname",
	Role:       "role",
	Solved:     "solved",
	Submitted:  "submitted",
	Rating:     "rating",
	CfId:       "cf_id",
	AtcId:      "atc_id",
	Company:    "company",
	Department: "department",
	Major:      "major",
	Class:      "class",
	Email:      "email",
	Tel:        "tel",
	Avatar:     "avatar",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
	DeleteAt:   "delete_at",
}

// NewUserBaseDao creates and returns a new DAO object for table data access.
func NewUserBaseDao(handlers ...gdb.ModelHandler) *UserBaseDao {
	return &UserBaseDao{
		group:    "default",
		table:    "user_base",
		columns:  userBaseColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserBaseDao) Columns() UserBaseColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserBaseDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysOperationLog = "sys_operation_log"

// SysOperationLog mapped from table <sys_operation_log>
type SysOperationLog struct {
	ID        int64      `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id,string"`
	UserID    int64      `gorm:"column:user_id;type:int unsigned;not null" json:"user_id"`
	Path      string     `gorm:"column:path;type:varchar(255);not null" json:"path"`
	Method    string     `gorm:"column:method;type:varchar(10);not null" json:"method"`
	IP        string     `gorm:"column:ip;type:varchar(15);not null" json:"ip"`
	Input     string     `gorm:"column:input;type:text;not null" json:"input"`
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName SysOperationLog's table name
func (*SysOperationLog) TableName() string {
	return TableNameSysOperationLog
}

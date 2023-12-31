// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysRoleResourceRel = "sys_role_resource_rel"

// SysRoleResourceRel mapped from table <sys_role_resource_rel>
type SysRoleResourceRel struct {
	ID         int64     `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id,string"`
	RoleID     int64     `gorm:"column:role_id;type:int;not null" json:"role_id"`
	ResourceID int64     `gorm:"column:resource_id;type:int;not null" json:"resource_id"`
	Created    time.Time `gorm:"column:created;type:datetime;not null" json:"created"`
}

// TableName SysRoleResourceRel's table name
func (*SysRoleResourceRel) TableName() string {
	return TableNameSysRoleResourceRel
}

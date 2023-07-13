// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysLogintrace = "sys_logintrace"

// SysLogintrace mapped from table <sys_logintrace>
type SysLogintrace struct {
	ID         int64      `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:序号" json:"id,string"` // 序号
	User       *string    `gorm:"column:user;type:varchar(30);comment:用户名" json:"user"`                            // 用户名
	RemoteAddr *string    `gorm:"column:remoteAddr;type:varchar(50);comment:IP地址" json:"remoteAddr"`               // IP地址
	LoginTime  *time.Time `gorm:"column:loginTime;type:datetime;comment:登录时间" json:"loginTime"`                    // 登录时间
}

// TableName SysLogintrace's table name
func (*SysLogintrace) TableName() string {
	return TableNameSysLogintrace
}

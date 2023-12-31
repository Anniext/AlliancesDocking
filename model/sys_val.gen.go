// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysVal = "sys_val"

// SysVal mapped from table <sys_val>
type SysVal struct {
	ID         int64      `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:序号" json:"id,string"`             // 序号
	Code       *string    `gorm:"column:code;type:varchar(20);comment:代码" json:"code"`                                      // 代码
	Desc       *string    `gorm:"column:desc;type:varchar(60);comment:描述" json:"desc"`                                      // 描述
	Value      *string    `gorm:"column:value;type:varchar(20);comment:变量" json:"value"`                                    // 变量
	Uplimit    *string    `gorm:"column:uplimit;type:varchar(20);comment:上限" json:"uplimit"`                                // 上限
	Step       *string    `gorm:"column:step;type:varchar(10);comment:步长" json:"step"`                                      // 步长
	Tag        *int64     `gorm:"column:tag;type:int;comment:状态:0.正常1.删除" json:"tag"`                                       // 状态:0.正常1.删除
	Createuser *string    `gorm:"column:createuser;type:varchar(20);comment:创建人" json:"createuser"`                         // 创建人
	Createdate *time.Time `gorm:"column:createdate;type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createdate"` // 创建时间
	Changeuser *string    `gorm:"column:changeuser;type:varchar(20);comment:修改人" json:"changeuser"`                         // 修改人
	Changedate *time.Time `gorm:"column:changedate;type:datetime;default:CURRENT_TIMESTAMP;comment:修改时间" json:"changedate"` // 修改时间
}

// TableName SysVal's table name
func (*SysVal) TableName() string {
	return TableNameSysVal
}

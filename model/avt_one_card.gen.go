// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAvtOneCard = "avt_one_card"

// AvtOneCard mapped from table <avt_one_card>
type AvtOneCard struct {
	ID           *int64     `gorm:"column:id;type:int" json:"id,string"`
	CardNo       *string    `gorm:"column:card_no;type:varchar(50)" json:"card_no"`
	Name         *string    `gorm:"column:name;type:varchar(50)" json:"name"`
	College      *string    `gorm:"column:college;type:varchar(50)" json:"college"`
	Status       *int64     `gorm:"column:status;type:int" json:"status"`
	CreateTime   *time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
	JobNumber    string     `gorm:"column:job_number;type:varchar(40);primaryKey" json:"job_number"`
	CardType     *int64     `gorm:"column:card_type;type:int" json:"card_type"`
	Subject      *string    `gorm:"column:subject;type:varchar(50)" json:"subject"`
	Class        *string    `gorm:"column:class;type:varchar(50)" json:"class"`
	CardUserType *int64     `gorm:"column:card_user_type;type:int" json:"card_user_type"`
	CollegeID    *int64     `gorm:"column:college_id;type:int;comment:学院编号" json:"college_id"` // 学院编号
	ClassID      *int64     `gorm:"column:class_id;type:int;comment:班级编号" json:"class_id"`     // 班级编号
}

// TableName AvtOneCard's table name
func (*AvtOneCard) TableName() string {
	return TableNameAvtOneCard
}

package data

import "AlliancesDocking/model"

type AvtWorkOrder struct {
	model.AvtWorkOrder
	Building string `orm:"-" json:"building"`
	Floor    string `orm:"-" json:"floor"`
	Room     string `orm:"-" json:"room"`
}

package models

import (
	"go-admin/common/models"
)

type RecordNode struct {
	models.Model

	Address string `json:"address" gorm:"type:varchar(128);comment:连接地址"`
	Key     string `json:"key" gorm:"type:varchar(255);comment:apikey"`
	Max     int64  `json:"max" gorm:"type:int(11);comment:最大任务数量"`
	Task    int64  `json:"task" gorm:"type:int(11);comment:当前任务数量"`
	Comment string `json:"comment" gorm:"type:varchar(128);comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (RecordNode) TableName() string {
	return "record_node"
}

func (e *RecordNode) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *RecordNode) GetId() interface{} {
	return e.Id
}

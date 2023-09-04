package models

import (
	"go-admin/common/models"
)

type Record struct {
	models.Model

	RoomId       int64  `json:"roomId" gorm:"type:int(11);comment:直播间ID"`
	StreamerId   int64  `json:"streamerId" gorm:"type:int(11);comment:主播ID"`
	StreamerName string `json:"streamerName" gorm:"type:varchar(128);comment:主播名称"`
	DeptId       int64  `json:"deptId" gorm:"type:int(11);comment:组ID"`
	NodeId       int64  `json:"nodeId" gorm:"type:int(11);comment:使用节点ID"`
	models.ModelTime
	models.ControlBy
}

func (Record) TableName() string {
	return "record"
}

func (e *Record) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Record) GetId() interface{} {
	return e.Id
}

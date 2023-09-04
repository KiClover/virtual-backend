package dto

import (
	"go-admin/app/record/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RecordGetPageReq struct {
	dto.Pagination `search:"-"`
	RoomId         int64  `form:"roomId"  search:"type:exact;column:room_id;table:record" comment:"直播间ID"`
	StreamerId     int64  `form:"streamerId"  search:"type:exact;column:streamer_id;table:record" comment:"主播ID"`
	StreamerName   string `form:"streamerName"  search:"type:exact;column:streamer_name;table:record" comment:"主播名称"`
	DeptId         int64  `form:"deptId"  search:"type:exact;column:dept_id;table:record" comment:"组ID"`
	NodeId         int64  `form:"nodeId"  search:"type:exact;column:node_id;table:record" comment:"使用节点ID"`
	RecordOrder
}

type RecordOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:record"`
	RoomId       string `form:"roomIdOrder"  search:"type:order;column:room_id;table:record"`
	StreamerId   string `form:"streamerIdOrder"  search:"type:order;column:streamer_id;table:record"`
	StreamerName string `form:"streamerNameOrder"  search:"type:order;column:streamer_name;table:record"`
	DeptId       string `form:"deptIdOrder"  search:"type:order;column:dept_id;table:record"`
	NodeId       string `form:"nodeIdOrder"  search:"type:order;column:node_id;table:record"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:record"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:record"`
	DeletedAt    string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:record"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:record"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:record"`
}

func (m *RecordGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RecordInsertReq struct {
	Id           int    `json:"-" comment:"编码"` // 编码
	RoomId       int64  `json:"roomId" comment:"直播间ID"`
	StreamerId   int64  `json:"streamerId" comment:"主播ID"`
	StreamerName string `json:"streamerName" comment:"主播名称"`
	DeptId       int64  `json:"deptId" comment:"组ID"`
	NodeId       int64  `json:"nodeId" comment:"使用节点ID"`
	common.ControlBy
}

func (s *RecordInsertReq) Generate(model *models.Record) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.RoomId = s.RoomId
	model.StreamerId = s.StreamerId
	model.StreamerName = s.StreamerName
	model.DeptId = s.DeptId
	model.NodeId = s.NodeId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *RecordInsertReq) GetId() interface{} {
	return s.Id
}

type RecordUpdateReq struct {
	Id           int    `uri:"id" comment:"编码"` // 编码
	RoomId       int64  `json:"roomId" comment:"直播间ID"`
	StreamerId   int64  `json:"streamerId" comment:"主播ID"`
	StreamerName string `json:"streamerName" comment:"主播名称"`
	DeptId       int64  `json:"deptId" comment:"组ID"`
	NodeId       int64  `json:"nodeId" comment:"使用节点ID"`
	common.ControlBy
}

func (s *RecordUpdateReq) Generate(model *models.Record) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.RoomId = s.RoomId
	model.StreamerId = s.StreamerId
	model.StreamerName = s.StreamerName
	model.DeptId = s.DeptId
	model.NodeId = s.NodeId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *RecordUpdateReq) GetId() interface{} {
	return s.Id
}

// RecordGetReq 功能获取请求参数
type RecordGetReq struct {
	Id int `uri:"id"`
}

func (s *RecordGetReq) GetId() interface{} {
	return s.Id
}

// RecordDeleteReq 功能删除请求参数
type RecordDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RecordDeleteReq) GetId() interface{} {
	return s.Ids
}

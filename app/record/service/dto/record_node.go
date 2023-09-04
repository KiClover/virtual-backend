package dto

import (
	"go-admin/app/record/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type RecordNodeGetPageReq struct {
	dto.Pagination `search:"-"`
	Address        string `form:"address"  search:"type:exact;column:address;table:record_node" comment:"连接地址"`
	Comment        string `form:"comment"  search:"type:exact;column:comment;table:record_node" comment:"备注"`
	RecordNodeOrder
}

type RecordNodeOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:record_node"`
	Address   string `form:"addressOrder"  search:"type:order;column:address;table:record_node"`
	Key       string `form:"keyOrder"  search:"type:order;column:key;table:record_node"`
	Max       string `form:"maxOrder"  search:"type:order;column:max;table:record_node"`
	Task      string `form:"taskOrder"  search:"type:order;column:task;table:record_node"`
	Comment   string `form:"commentOrder"  search:"type:order;column:comment;table:record_node"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:record_node"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:record_node"`
	DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:record_node"`
	CreateBy  string `form:"createByOrder"  search:"type:order;column:create_by;table:record_node"`
	UpdateBy  string `form:"updateByOrder"  search:"type:order;column:update_by;table:record_node"`
}

func (m *RecordNodeGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type RecordNodeInsertReq struct {
	Id      int    `json:"-" comment:"编码"` // 编码
	Address string `json:"address" comment:"连接地址"`
	Key     string `json:"key" comment:"apikey"`
	Max     int64  `json:"max" comment:"最大任务数量"`
	Task    int64  `json:"task" comment:"当前任务数量"`
	Comment string `json:"comment" comment:"备注"`
	common.ControlBy
}

func (s *RecordNodeInsertReq) Generate(model *models.RecordNode) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Address = s.Address
	model.Key = s.Key
	model.Max = s.Max
	model.Task = s.Task
	model.Comment = s.Comment
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *RecordNodeInsertReq) GetId() interface{} {
	return s.Id
}

type RecordNodeUpdateReq struct {
	Id      int    `uri:"id" comment:"编码"` // 编码
	Address string `json:"address" comment:"连接地址"`
	Key     string `json:"key" comment:"apikey"`
	Max     int64  `json:"max" comment:"最大任务数量"`
	Task    int64  `json:"task" comment:"当前任务数量"`
	Comment string `json:"comment" comment:"备注"`
	common.ControlBy
}

func (s *RecordNodeUpdateReq) Generate(model *models.RecordNode) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Address = s.Address
	model.Key = s.Key
	model.Max = s.Max
	model.Task = s.Task
	model.Comment = s.Comment
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *RecordNodeUpdateReq) GetId() interface{} {
	return s.Id
}

// RecordNodeGetReq 功能获取请求参数
type RecordNodeGetReq struct {
	Id int `uri:"id"`
}

func (s *RecordNodeGetReq) GetId() interface{} {
	return s.Id
}

// RecordNodeDeleteReq 功能删除请求参数
type RecordNodeDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *RecordNodeDeleteReq) GetId() interface{} {
	return s.Ids
}

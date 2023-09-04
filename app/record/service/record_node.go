package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/record/models"
	"go-admin/app/record/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type RecordNode struct {
	service.Service
}

// GetPage 获取RecordNode列表
func (e *RecordNode) GetPage(c *dto.RecordNodeGetPageReq, p *actions.DataPermission, list *[]models.RecordNode, count *int64) error {
	var err error
	var data models.RecordNode

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RecordNodeService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取RecordNode对象
func (e *RecordNode) Get(d *dto.RecordNodeGetReq, p *actions.DataPermission, model *models.RecordNode) error {
	var data models.RecordNode

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRecordNode error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}
func (e *RecordNode) GetNodeConfig(nodeid int64, model *models.RecordNode) error {
	var data models.RecordNode

	err := e.Orm.Model(&data).
		First(model, nodeid).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRecordNode error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建RecordNode对象
func (e *RecordNode) Insert(c *dto.RecordNodeInsertReq) error {
	var err error
	var data models.RecordNode
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RecordNodeService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改RecordNode对象
func (e *RecordNode) Update(c *dto.RecordNodeUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.RecordNode{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RecordNodeService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除RecordNode
func (e *RecordNode) Remove(d *dto.RecordNodeDeleteReq, p *actions.DataPermission) error {
	var data models.RecordNode

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRecordNode error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

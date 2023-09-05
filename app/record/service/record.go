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

type Record struct {
	service.Service
}

// GetPage 获取Record列表
func (e *Record) GetPage(c *dto.RecordGetPageReq, p *actions.DataPermission, list *[]models.Record, count *int64) error {
	var err error
	var data models.Record

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("RecordService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Record对象
func (e *Record) Get(d *dto.RecordGetReq, p *actions.DataPermission, model *models.Record) error {
	var data models.Record

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetRecord error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

func (e *Record) GetRecordByDept(p *actions.DataPermission, model *models.Record) (int, error) {
	var data models.Record

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return 1, nil
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return 0, err
	}
	return 0, nil
}

// Insert 创建Record对象
func (e *Record) Insert(c *dto.RecordInsertReq) error {
	var err error
	var data models.Record
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RecordService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

func (e *Record) InsertTask(p *actions.DataPermission, c *dto.RecordInsertTaskReq, node int64) error {
	var err error
	var data models.Record
	c.GenerateTask(&data)
	var i int64
	err = e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		Count(&i).Error
	if i > 0 {
		err = errors.New("数据库已经记录录播任务，不可创建")
		return err
	}
	data.NodeId = node
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("RecordService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Record对象
func (e *Record) Update(c *dto.RecordUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Record{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RecordService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

func (e *Record) UpdateByDept(c *dto.RecordUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Record{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data)
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("RecordService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Record
func (e *Record) Remove(d *dto.RecordDeleteReq, p *actions.DataPermission) error {
	var data models.Record

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveRecord error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

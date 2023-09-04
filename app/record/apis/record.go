package apis

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/record/models"
	"go-admin/app/record/service"
	"go-admin/app/record/service/dto"
	"go-admin/common/actions"
	"strconv"
)

type Record struct {
	api.Api
}

// GetPage 获取录播任务管理列表
// @Summary 获取录播任务管理列表
// @Description 获取录播任务管理列表
// @Tags 录播任务管理
// @Param roomId query int64 false "直播间ID"
// @Param streamerId query int64 false "主播ID"
// @Param streamerName query string false "主播名称"
// @Param deptId query int64 false "组ID"
// @Param nodeId query int64 false "使用节点ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Record}} "{"code": 200, "data": [...]}"
// @Router /api/v1/record [get]
// @Security Bearer
func (e Record) GetPage(c *gin.Context) {
	req := dto.RecordGetPageReq{}
	s := service.Record{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Record, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取录播任务管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取录播任务管理
// @Summary 获取录播任务管理
// @Description 获取录播任务管理
// @Tags 录播任务管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Record} "{"code": 200, "data": [...]}"
// @Router /api/v1/record/{id} [get]
// @Security Bearer
func (e Record) Get(c *gin.Context) {
	req := dto.RecordGetReq{}
	s := service.Record{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Record

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取录播任务管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建录播任务管理
// @Summary 创建录播任务管理
// @Description 创建录播任务管理
// @Tags 录播任务管理
// @Accept application/json
// @Product application/json
// @Param data body dto.RecordInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/record [post]
// @Security Bearer
func (e Record) Insert(c *gin.Context) {
	req := dto.RecordInsertReq{}
	s := service.Record{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建录播任务管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改录播任务管理
// @Summary 修改录播任务管理
// @Description 修改录播任务管理
// @Tags 录播任务管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RecordUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/record/{id} [put]
// @Security Bearer
func (e Record) Update(c *gin.Context) {
	req := dto.RecordUpdateReq{}
	s := service.Record{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改录播任务管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除录播任务管理
// @Summary 删除录播任务管理
// @Description 删除录播任务管理
// @Tags 录播任务管理
// @Param data body dto.RecordDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/record [delete]
// @Security Bearer
func (e Record) Delete(c *gin.Context) {
	s := service.Record{}
	req := dto.RecordDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除录播任务管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// GetStatus Get 获取录播任务状态
// @Summary 获取录播任务状态
// @Description 获取录播任务状态
// @Tags 录播任务状态
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/record/status [get]
// @Security Bearer
func (e Record) GetStatus(c *gin.Context) {
	sr := service.Record{}
	sn := service.RecordNode{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&sr.Service).
		MakeService(&sn.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Record
	var objectnode models.RecordNode
	p := actions.GetPermissionFromContext(c)
	noexist, err := sr.GetRecordByDept(p, &object)
	// 验证数据库内录播任务是否存在
	if noexist == 1 {
		err = errors.New("数据库中未记录录播任务")
		e.Error(901, err, "任务错误，无法获取数据库中的任务详情")
		return
	}
	// 获取节点信息
	err = sn.GetNodeConfig(object.NodeId, &objectnode)
	if err != nil {
		e.Error(902, err, fmt.Sprintf("获取录播节点信息管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	roomid := strconv.FormatInt(object.RoomId, 10)
	// 向节点后端发送请求
	resp, err := actions.RecRequest("GET",
		objectnode.Address+"/api/v1/tasks/"+roomid+"/data",
		objectnode.Key, nil)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("向录播节点发送请求失败，\r\n失败信息 %s", err.Error()))
		return
	}
	if resp.RoomInfo.RoomID != int(object.RoomId) {
		e.Error(903, nil, "未向边缘录播节点取得录播任务数据")
		return
	}
	e.OK(resp, "获取成功")
}

func (e Record) InsertTask(c *gin.Context) {
	req := dto.RecordInsertReq{}
	s := service.Record{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建录播任务管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

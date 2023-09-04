package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/record/models"
	"go-admin/app/record/service"
	"go-admin/app/record/service/dto"
	"go-admin/common/actions"
)

type RecordNode struct {
	api.Api
}

// GetPage 获取服务节点列表
// @Summary 获取服务节点列表
// @Description 获取服务节点列表
// @Tags 服务节点
// @Param address query string false "连接地址"
// @Param comment query string false "备注"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.RecordNode}} "{"code": 200, "data": [...]}"
// @Router /api/v1/record-node [get]
// @Security Bearer
func (e RecordNode) GetPage(c *gin.Context) {
	req := dto.RecordNodeGetPageReq{}
	s := service.RecordNode{}
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
	list := make([]models.RecordNode, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取服务节点失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取服务节点
// @Summary 获取服务节点
// @Description 获取服务节点
// @Tags 服务节点
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.RecordNode} "{"code": 200, "data": [...]}"
// @Router /api/v1/record-node/{id} [get]
// @Security Bearer
func (e RecordNode) Get(c *gin.Context) {
	req := dto.RecordNodeGetReq{}
	s := service.RecordNode{}
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
	var object models.RecordNode

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取服务节点失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建服务节点
// @Summary 创建服务节点
// @Description 创建服务节点
// @Tags 服务节点
// @Accept application/json
// @Product application/json
// @Param data body dto.RecordNodeInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/record-node [post]
// @Security Bearer
func (e RecordNode) Insert(c *gin.Context) {
	req := dto.RecordNodeInsertReq{}
	s := service.RecordNode{}
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
		e.Error(500, err, fmt.Sprintf("创建服务节点失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改服务节点
// @Summary 修改服务节点
// @Description 修改服务节点
// @Tags 服务节点
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.RecordNodeUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/record-node/{id} [put]
// @Security Bearer
func (e RecordNode) Update(c *gin.Context) {
	req := dto.RecordNodeUpdateReq{}
	s := service.RecordNode{}
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
		e.Error(500, err, fmt.Sprintf("修改服务节点失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除服务节点
// @Summary 删除服务节点
// @Description 删除服务节点
// @Tags 服务节点
// @Param data body dto.RecordNodeDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/record-node [delete]
// @Security Bearer
func (e RecordNode) Delete(c *gin.Context) {
	s := service.RecordNode{}
	req := dto.RecordNodeDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除服务节点失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

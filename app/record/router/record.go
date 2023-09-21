package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/record/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerRecordRouter)
}

// registerRecordRouter
func registerRecordRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Record{}
	r := v1.Group("/record").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		// 基础CRUD部分
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
		// 任务信息处理部分
		r.GET("/status", actions.PermissionAction(), api.GetStatus)
		r.GET("/config", actions.PermissionAction(), api.GetConfig)
		r.POST("/creat", actions.PermissionAction(), api.InsertTask)
		r.PUT("/update", actions.PermissionAction(), api.UpdateTask)
		r.DELETE("/", actions.PermissionAction(), api.DelTask)
		// 任务处理部分
		r.GET("/start", actions.PermissionAction(), api.StartTask)
		r.GET("/stop", actions.PermissionAction(), api.StopTask)
		r.GET("/cut", actions.PermissionAction(), api.CutTask)
	}
}

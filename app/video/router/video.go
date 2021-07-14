package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/video/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, videoCheckRoleRouter)
}

func videoCheckRoleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r := v1.Group("/video").Use(authMiddleware.MiddlewareFunc())
	{
		video := apis.Video{}
		r.POST("/applyUpload", video.ApplyUpload)
		//r.GET("/listVideos", gen.GenCode)
		//r.GET("/getVideo", gen.GenApiToFile)
		//r.GET("/updateVideoStatus", gen.GenMenuAndApi)
		//r.GET("addOrUpdateVideo", sysTable.GetSysTablesTree)
	}
}

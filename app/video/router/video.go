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
	//r := v1.Group("/video").Use(authMiddleware.MiddlewareFunc())  // 开启权限校验
	r := v1.Group("/video")
	{
		video := apis.Video{}
		r.POST("/applyUpload", video.ApplyUpload)
		r.POST("/listVideos", video.ListVideos)
		r.POST("/getVideo", video.GetVideo)
		r.POST("/updateVideoStatus", video.UpdateVideoStatus)
		r.POST("/addOrUpdateVideo", video.AddOrUpdateVideo)
	}
}

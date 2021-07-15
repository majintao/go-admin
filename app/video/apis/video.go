package apis

import (
	"context"
	"encoding/json"
	"git.nonolive.co/FE/mildom-video/pkg/app/mildomapi"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/rpc"
)

type Video struct {
	api.Api
}

// ApplyUpload
// @Summary 申请上传凭证
// @Description 后台获取申请上传的凭证
// @Tags 视频
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/video/applyUpload [post]
// @Security Bearer
func (v Video) ApplyUpload(c *gin.Context) {
	// 申请上传接口
	applyUploadReq := &mildomapi.ApplyUploadReq{}
	applyUploadReq.Origin = 1
	upload, err := rpc.MildomVideoServiceClient.ApplyUpload(context.Background(), applyUploadReq)
	if err != nil {
		v.Logger.Error(err)
		v.Error(500, err, "申请token失败")
		return
	}

	result := make(map[string]interface{})
	json.Unmarshal(upload.UploadParams, &result)
	v.OK(result, "申请成功")
}

// ListVideos
// @Summary 根据条件查询视频列表
// @Description 根据条件查询获取视频列表
// @Tags 视频
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserControl true "用户数据"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/video/listVideos [post]
// @Security Bearer
func (v Video) ListVideos(c *gin.Context) {

}

// GetVideo
// @Summary 根据视频id获取视频信息
// @Description 根据视频id获取视频信息
// @Tags 视频
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserControl true "用户数据"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/video/getVideo [post]
// @Security Bearer
func (v Video) GetVideo(c *gin.Context) {

}

// UpdateVideoStatus
// @Summary 根据视频id更新视频状态
// @Description 根据视频id更新视频状态
// @Tags 视频
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserControl true "用户数据"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/video/updateVideoStatus [post]
// @Security Bearer
func (v Video) UpdateVideoStatus(c *gin.Context) {

}

// AddOrUpdateVideo
// @Summary 新增或者修改视频
// @Description 新增或者修改视频
// @Tags 视频
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysUserControl true "用户数据"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/video/AddOrUpdateVideo [post]
// @Security Bearer
func (v Video) AddOrUpdateVideo(c *gin.Context) {

}

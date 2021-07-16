package apis

import (
	"context"
	"encoding/json"
	"git.nonolive.co/FE/mildom-video/pkg/app/dto/cms"
	"git.nonolive.co/FE/mildom-video/pkg/app/mildomapi"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/rpc"
	"go-admin/app/video/dto"
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
	v.MakeContext(c)
	applyUploadReq := &mildomapi.ApplyUploadReq{}
	applyUploadReq.Origin = 1
	upload, err := rpc.MildomVideoCmsServiceClient.ApplyUpload(context.Background(), applyUploadReq)
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
// @Param data body cms.ListVideosReqDto true "视频列表选择参数"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/video/listVideos [post]
// @Security Bearer
func (v Video) ListVideos(c *gin.Context) {
	listVideosReqDto := cms.ListVideosReqDto{}
	v.MakeContext(c).Bind(&listVideosReqDto)

	marshal, err := json.Marshal(listVideosReqDto)
	if err != nil {
		return
	}
	req := &mildomapi.ListVideosReq{
		Params: marshal,
	}
	result, err := rpc.MildomVideoCmsServiceClient.ListVideos(context.Background(), req)
	if err != nil {
		v.Logger.Error(err)
		v.Error(500, err, "获取列表失败")
		return
	}

	dto := cms.ListVideosRespDto{}
	json.Unmarshal(result.Results, &dto)

	v.OK(dto, "请求成功")
}

// GetVideo
// @Summary 根据视频id获取视频信息
// @Description 根据视频id获取视频信息
// @Tags 视频
// @Accept  application/json
// @Product application/json
// @Param data body dto.GetVideoDto true "获取视频信息查询id"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/video/getVideo [post]
// @Security Bearer
func (v Video) GetVideo(c *gin.Context) {
	videoDto := dto.GetVideoDto{}
	v.MakeContext(c).Bind(&videoDto)

	g := &mildomapi.GetVideoReq{
		Id: videoDto.Id,
	}

	resp, err := rpc.MildomVideoCmsServiceClient.GetVideo(context.Background(), g)
	if err != nil {
		v.Logger.Error(err)
		v.Error(500, err, "获取视频信息失败")
		return
	}

	dto := cms.ListVideosRespDto{}
	json.Unmarshal(resp.Results, &dto)

	v.OK(dto, "请求成功")
}

// UpdateVideoStatus
// @Summary 根据视频id更新视频状态
// @Description 根据视频id更新视频状态
// @Tags 视频
// @Accept  application/json
// @Product application/json
// @Param data body cms.UpdateStatusReqDto true "更新视频状态请求"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/video/updateVideoStatus [post]
// @Security Bearer
func (v Video) UpdateVideoStatus(c *gin.Context) {
	reqDto := cms.UpdateStatusReqDto{}
	v.MakeContext(c).Bind(&reqDto)

	marshal, err := json.Marshal(reqDto)
	if err != nil {
		v.Logger.Error(err)
		v.Error(500, err, "更新失败")
		return
	}

	req := &mildomapi.UpdateVideoStatusReq{
		Params: marshal,
	}
	_, err = rpc.MildomVideoCmsServiceClient.UpdateVideoStatus(context.Background(), req)
	if err != nil {
		v.Logger.Error(err)
		v.Error(500, err, "更新失败")
		return
	}

	v.OK(nil, "更新成功")
}

// AddOrUpdateVideo
// @Summary 新增或者修改视频
// @Description 新增或者修改视频
// @Tags 视频
// @Accept  application/json
// @Product application/json
// @Param data body cms.VideoDto true "用户数据"
// @Success 200 {string} {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/video/addOrUpdateVideo [post]
// @Security Bearer
func (v Video) AddOrUpdateVideo(c *gin.Context) {
	videoDto := cms.VideoDto{}
	v.MakeContext(c).Bind(&videoDto)

	marshal, _ := json.Marshal(videoDto)
	if videoDto.Id == "" {
		// 新增
		req := &mildomapi.SaveVideoReq{}
		req.Params = marshal
		resp, err := rpc.MildomVideoCmsServiceClient.SaveVideo(context.Background(), req)
		if err != nil {
			v.Logger.Error(err)
			v.Error(500, err, "更新失败")
			return
		}

		m := make(map[string]interface{})
		m["id"] = resp.Id
		v.OK(m, "操作成功")
	} else {
		// 修改
		req := &mildomapi.UpdateVideoReq{}
		req.Params = marshal
		_, err := rpc.MildomVideoCmsServiceClient.UpdateVideo(context.Background(), req)
		if err != nil {
			v.Logger.Error(err)
			v.Error(500, err, "更新失败")
			return
		}

		v.OK(nil, "操作成功")
	}

}

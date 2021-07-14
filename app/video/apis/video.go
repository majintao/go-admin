package apis

import (
	"context"
	"git.nonolive.co/FE/mildom-video/pkg/app/mildomapi"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/app/video/rpc"
)

type Video struct {
	api.Api
}

func (v Video) ApplyUpload(c *gin.Context) {
	s := service.SysUser{}
	req := dto.SysUserControl{}
	err := v.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		v.Logger.Error(err)
		v.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		v.Logger.Error(err)
		v.Error(500, err, err.Error())
		return
	}

	// 申请上传接口
	applyUploadReq := &mildomapi.ApplyUploadReq{}
	rpc.MildomVideoServiceClient.ApplyUpload(context.Background(), applyUploadReq)
	v.OK(req.GetId(), "创建成功")
}

func (v Video) ListVideos(c *gin.Context) {

}

func (v Video) GetVideo(c *gin.Context) {

}

func (v Video) UpdateVideoStatus(c *gin.Context) {

}

func (v Video) AddOrUpdateVideo(c *gin.Context) {

}

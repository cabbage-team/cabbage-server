package controller

import (
	"cabbage-server/internal"
	"cabbage-server/response"
	"cabbage-server/service"
	"crypto/sha256"
	"encoding/hex"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// PostSearch
// @Summary profile share
// @Description 个人主页分享
// @Tags profile
// @Param name path string true "the user name"
// @Accept json
// @Router /v1/api/bio/{name} [get]
func ProfileSharre(ctx *gin.Context) {
	name := ctx.Param("name")
	profile, err := service.ProfileShare(name)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, gin.H{
		"data": profile,
	})
}

// PostSearch
// @Summary uploading user avatar
// @Description 提交个人头像
// @Tags profile
// @Accept json
// @Param file formData file true "用户头像"
// @Router /v1/api/profile/avatar [post]
func UploadingAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {

	}
	h := sha256.New()
	switch filepath.Ext(file.Filename) {
	case ".png", "jpg", ".webp",".gif",".jpeg":
		h.Write([]byte(file.Filename + "cabbage"))
		newFilename := hex.EncodeToString(h.Sum(nil))

		newFilenameWithExt := newFilename + filepath.Ext(file.Filename)

		c.SaveUploadedFile(file, "static/"+newFilenameWithExt)
	default:
		response.Error(c,internal.UnsupportedFileType)
	}
}

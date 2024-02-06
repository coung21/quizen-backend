package transport

import (
	"net/http"
	"quizen/common"
	"quizen/module/upload/model"

	"github.com/gin-gonic/gin"
)

// @Summary Upload file
// @Description Upload file to S3
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Param folder formData string false "Folder to save the file"
// @Success 200 {object} common.Image
// @Failure 400 {object} common.ErrResp
// @Router /upload [post]
func (h *httpHandler) UploadHdl() gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, "File is required", nil))
			return
		}

		folder := c.DefaultPostForm("folder", "images")

		file, err := fileHeader.Open()

		defer file.Close()

		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, common.BadRequest.Error(), nil))
			return
		}

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, common.BadRequest.Error(), nil))
			return
		}

		img, err := h.uc.Upload(c.Request.Context(), dataBytes, fileHeader.Filename, folder)

		if err != nil {
			if err == model.ErrFileIsNotImage || err == model.ErrFileTooLarge {
				c.JSON(http.StatusBadRequest, common.NewRestErr(http.StatusBadRequest, err.Error(), err))
				return
			} else {
				c.JSON(http.StatusInternalServerError, common.NewRestErr(http.StatusInternalServerError, err.Error(), err))
				return
			}
		}

		c.JSON(http.StatusOK, img)
	}
}

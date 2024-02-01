package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20

	router.POST("/images", func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		savepath := filepath.Join("images", file.Filename)
		if err := c.SaveUploadedFile(file, savepath); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		c.String(http.StatusOK, "File %s uploaded successfully.", file.Filename)
	})

	router.Run(":8080")
}

package main

import (
	"mime/multipart"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct{}
type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func (handler *Handler) HandleParseVideo(ctx *gin.Context) {
	var form Form
	err := ctx.ShouldBind(&form)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(403, map[string]string{"error": err.Error()})
	}
	parsedFilename := strings.Split(form.File.Filename, ".mp4")[0]
	zipPath := fmt.Sprintf("./tmp/%s", parsedFilename)
	// Cleanup(zipPath)
	fmt.Println("Saving video to /tmp/")
	err = ctx.SaveUploadedFile(form.File, "./tmp/vid.mp4")
	if err != nil {
		ctx.JSON(403, map[string]string{"error": err.Error()})
	}
	fmt.Println("Parsing " + parsedFilename)
	parse("./tmp/vid.mp4", parsedFilename)
	if err := zipSource(zipPath, "output.zip"); err != nil {
		ctx.JSON(403, map[string]string{"error": err.Error()})
	}
	
	if err != nil {
		ctx.JSON(403, map[string]string{"error": err.Error()})
	}
	Cleanup(zipPath)
	ctx.File("./output.zip")
}
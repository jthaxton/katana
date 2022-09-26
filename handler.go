package main

import (
	"mime/multipart"
	"fmt"
	"strings"
	"os"
	"io"
	"net/url"

	"github.com/kkdai/youtube/v2"
	"github.com/gin-gonic/gin"
)

type Handler struct{}
type Form struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func (handler *Handler) HandleGetVideo(ctx *gin.Context) {
	u := ctx.Query("name")
	if len(u) == 0 {
		ctx.JSON(403,map[string]string{"error":"youtubeId not found."})
		return
	}
	parsed, err := url.Parse(u)
	if err != nil {
		ctx.JSON(403,map[string]string{"error":err.Error()})
		return
	}

	query := parsed.Query()
	videoID := query.Get("v")
	if len(videoID) == 0 {
		ctx.JSON(403,map[string]string{"error":"video id not found."})
		return
	}
	// videoID := "BaW_jenozKc"
	client := youtube.Client{}
	fmt.Println("ID is " + videoID)
	video, err := client.GetVideo(videoID)
	if err != nil {
		ctx.JSON(403,map[string]string{"error": err.Error()})
		return
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		ctx.JSON(403,map[string]string{"error": err.Error()})
		return
	}

	file, err := os.Create("./tmp/vid.mp4")
	if err != nil {
		ctx.JSON(403,map[string]string{"error": err.Error()})
		return
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		ctx.JSON(403,map[string]string{"error":err.Error()})
		return
	}
	parsedFilename := videoID

	fmt.Println("Parsing " + parsedFilename)
	zipPath := fmt.Sprintf("./tmp/%s", parsedFilename)

	parse("./tmp/vid.mp4", parsedFilename)
	if err := zipSource(zipPath, "output.zip"); err != nil {
		ctx.JSON(403, map[string]string{"error": err.Error()})
	}
	
	if err != nil {
		ctx.JSON(403, map[string]string{"error": err.Error()})
	}
	ctx.File("./output.zip")
	Cleanup(zipPath)
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
	ctx.File("./output.zip")
	Cleanup(zipPath)
}

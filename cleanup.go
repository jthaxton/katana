package main

import (
	"fmt"
	"os"
)

func Cleanup(path string) {
	fmtPath := fmt.Sprintf(path)
	fmt.Println("Removing temp files from:" + fmtPath)
	os.RemoveAll(fmtPath)
	os.RemoveAll("./tmp/vid.mp4")
	os.RemoveAll("./output.zip")
	
}
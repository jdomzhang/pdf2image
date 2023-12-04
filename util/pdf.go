package util

import (
	"fmt"
	"image/jpeg"
	"os"

	fitz "github.com/gen2brain/go-fitz"
)

// 从pdf文件提取所有页面，并转为图片
func ConvertPdfToImages(pdfPath string, imgFolder string) (imgPaths []string, err error) {
	doc, err := fitz.New(pdfPath)
	if err != nil {
		return
	}
	defer doc.Close()

	// 1. 创建目录
	if err = os.MkdirAll(imgFolder, os.ModePerm); err != nil {
		return
	}

	// 2. 创建文件
	totalPages := doc.NumPage()
	imgPaths = make([]string, totalPages)
	for i := 0; i < totalPages; i++ {
		imgPath := fmt.Sprintf("%s/%d.jpg", imgFolder, i+1)
		f, err := os.Create(imgPath)
		if err != nil {
			return imgPaths, err
		}

		// 3. 获取指定页的图片，保存到文件
		img, err := doc.Image(i)
		if err != nil {
			return imgPaths, err
		}
		err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
		if err != nil {
			return imgPaths, err
		}

		imgPaths[i] = imgPath
	}

	return
}

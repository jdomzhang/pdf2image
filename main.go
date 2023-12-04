package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jdomzhang/pdf2image/util"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "input",
				Value: "",
				Usage: "pdf文件路径",
			},
			&cli.StringFlag{
				Name:  "output",
				Value: "./output",
				Usage: "图片文件目录",
			},
		},
		Action: func(c *cli.Context) error {
			pdfFilePath := c.String("input")
			outputPath := c.String("output")
			if pdfFilePath == "" {
				return errors.New("请输入pdf文件路径")
			}
			if outputPath == "" {
				return errors.New("请输入图片文件目录")
			}

			// 校验文件是否存在
			if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
				return errors.New("pdf文件不存在")
			} else if err != nil {
				return err
			}

			// 校验目录是否存在
			if _, err := os.Stat(outputPath); os.IsNotExist(err) {
				// 创建目录
				if err := os.Mkdir(outputPath, os.ModePerm); err != nil {
					return err
				}
			} else if err != nil {
				return err
			}

			// pdf转成图片
			imgPaths, err := util.ConvertPdfToImages(pdfFilePath, outputPath)
			if err != nil {
				return err
			}

			for _, imgPath := range imgPaths {
				fmt.Println(imgPath)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

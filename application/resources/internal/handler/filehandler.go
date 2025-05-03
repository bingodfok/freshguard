package handler

import (
	"github.com/bingodfok/freshguard/application/resources/internal/alc"
	"github.com/bingodfok/freshguard/application/resources/internal/common/code"
	"github.com/bingodfok/freshguard/application/resources/internal/common/constant"
	"github.com/bingodfok/freshguard/application/resources/internal/logic"
	"github.com/bingodfok/freshguard/pkg/common/errorx"
	"github.com/bingodfok/freshguard/pkg/model/result"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"slices"
	"strings"
)

var ImgSuffix = []string{".jpg", ".jpeg", ".png", ".bmp", ".gif", ".ico"}

func NewImageHandler(appCtx *alc.ApplicationContext) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		form, err := ctx.MultipartForm()
		if err != nil {
			return err
		}
		files := form.File["file"]
		if len(files) == 0 {
			return errorx.NewCodeError(code.ParamError, "请上传正确的文件")
		}
		if len(files) > 1 {
			return errorx.NewCodeError(code.ParamError, "仅支持同时上传一个文件")
		}
		file := files[0]
		fileName := file.Filename
		index := strings.LastIndex(fileName, ".")
		if index != -1 && !slices.Contains(ImgSuffix, fileName[index:]) {
			return errorx.NewCodeError(code.ParamError, "不支持的文件类型")
		}
		open, err := file.Open()
		if err != nil {
			return errorx.NewDefaultCodeError("文件读取失败")
		}
		oss := logic.NewOssServiceLogic(appCtx)
		path := constant.ImageBasePath + uuid.New().String() + fileName[index:]
		uploaded, err := oss.ImageUpload(open, path)
		if err != nil {
			return errorx.NewDefaultCodeError("文件上传失败")
		}
		err = ctx.JSON(result.Success(map[string]interface{}{
			"url": uploaded,
		}))
		if err != nil {
			return err
		}
		return nil
	}

}

func TextHandler(appCtx *alc.ApplicationContext) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		err := ctx.SendString("Hello World")
		if err != nil {
			return err
		}
		return nil
	}

}

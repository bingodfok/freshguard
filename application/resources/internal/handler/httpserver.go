package handler

import (
	"errors"
	"github.com/bingodfok/freshguard/application/resources/internal/alc"
	"github.com/bingodfok/freshguard/pkg/common/errorx"
	"github.com/bingodfok/freshguard/pkg/model/result"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

var (
	OssRouterFunc = func(appCtx *alc.ApplicationContext) func(router fiber.Router) {
		return func(router fiber.Router) {
			router.Post("/imageUpload", NewImageHandler(appCtx))
			router.Get("/test", TextHandler(appCtx))
		}
	}
	BaseRouterFunc = func(appCtx *alc.ApplicationContext) func(router fiber.Router) {
		return func(router fiber.Router) {
			router.Route("/oss", OssRouterFunc(appCtx))
		}
	}
)

type HttpServer struct {
	appCtx *alc.ApplicationContext
}

func NewHttpServer(appCtx *alc.ApplicationContext) *HttpServer {
	return &HttpServer{
		appCtx: appCtx,
	}
}

func (server *HttpServer) Run() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       server.appCtx.Config.AppName,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			var codeError errorx.CodeError
			switch {
			case errors.As(err, &codeError):
				ctx.Status(fiber.StatusOK)
				err := ctx.JSON(result.Fail(codeError.Code, err.Error()))
				if err != nil {
					return err
				}
				break
			default:
				ctx.Status(fiber.StatusInternalServerError)
				err := ctx.JSON(result.Fail(fiber.StatusInternalServerError, "系统未知异常"))
				if err != nil {
					return err
				}
			}
			return nil
		},
	})
	app.Route("/api", BaseRouterFunc(server.appCtx))
	err := app.Listen(":" + strconv.Itoa(server.appCtx.Config.Web.Port))
	if err != nil {
		log.Fatal(err)
	}
}

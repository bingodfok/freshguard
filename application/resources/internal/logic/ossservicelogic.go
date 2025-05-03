package logic

import (
	"context"
	"github.com/bingodfok/freshguard/application/resources/internal/alc"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"io"
)

type OssServiceLogic struct {
	appCtx *alc.ApplicationContext
}

func NewOssServiceLogic(appCtx *alc.ApplicationContext) *OssServiceLogic {
	return &OssServiceLogic{appCtx: appCtx}
}

func (logic *OssServiceLogic) ImageUpload(reader io.Reader, fileName string) (string, error) {
	err := logic.appCtx.Uploader.UploadReader(context.Background(), reader, &uploader.ObjectOptions{
		BucketName: logic.appCtx.Config.Qiniu.PublicBucket,
		ObjectName: &fileName,
		FileName:   fileName,
	}, nil)
	if err != nil {
		return "", err
	}
	return logic.appCtx.Config.Qiniu.DomainUrl + "/" + fileName, nil
}

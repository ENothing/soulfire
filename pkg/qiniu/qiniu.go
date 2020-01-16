package qiniu

import (
	"context"
	"github.com/qiniu/api.v7/storage"
	"soulfire/pkg/qiniu/config"
)

func Upload(bucket, localFile string, key string) (string, error) {

	cfg := config.NewConfig()
	upToken := config.UpToken(bucket)

	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return "", err
	}

	return ret.Key, nil

}

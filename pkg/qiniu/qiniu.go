package qiniu

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/storage"
	"soulfire/pkg/qiniu/config"
	"soulfire/utils"
)

func Upload(bucket, localFile string) {

	key := utils.Uid("FE").png

	cfg := config.NewConfig()
	upToken := config.UpToken(bucket)

	formUploader := storage.NewFormUploader(cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)

}

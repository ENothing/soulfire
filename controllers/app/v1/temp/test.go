package temp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"soulfire/pkg/qiniu"
	"soulfire/pkg/rsp"
	"soulfire/utils"
)

func Upload(ctx *gin.Context)  {

	MediaUrl := "http://fxcgb.gongbo.vip"
	file, err := ctx.FormFile("file")
	fmt.Println(err)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.UploadErr, nil, "")
		return
	}
	bucket := "fxcgb"

	ext := path.Ext(file.Filename)
	key := utils.Uid("FE") + ext

	dst := "runtime/tmp/imgs/" + key

	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.UploadErr, nil, "")
		return
	}

	img, err := qiniu.Upload(bucket, dst, "image/"+key)

	url := MediaUrl + "/" + img

	if err != nil {
		rsp.JsonResonse(ctx, rsp.UploadErr, nil, "")
		return
	}

	_ = os.Remove(dst)

	rsp.JsonResonse(ctx, rsp.OK, url, "")

}
package user

import (
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"soulfire/models"
	"soulfire/pkg/config"
	"soulfire/pkg/qiniu"
	"soulfire/pkg/rsp"
	"soulfire/pkg/verify"
	"soulfire/utils"
)

type FeedbackForm struct {
	Title   string `json:"title" valid:"Required;" ch:"反馈标题"`
	Content string `json:"content" valid:"Required;" ch:"反馈内容"`
}

func PostFeedback(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(int64)
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	pics := ctx.PostForm("pics")


	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	feedbackForm := FeedbackForm{
		title,
		content,
	}

	message := verify.FormVerify(&feedbackForm)
	if message != nil {
		rsp.JsonResonse(ctx, rsp.FeedbackParamsEmpty, nil, message.(string))
		return
	}

	feedback := models.Feedback{
		UserId:  userId,
		Title:   title,
		Content: content,
		Pics:  pics ,
	}
	err := feedback.Create()

	if err != nil {
		rsp.JsonResonse(ctx, rsp.FeedbackPostFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func Upload(ctx *gin.Context) {

	app, _ := config.Cfg.GetSection("qiniu")
	MediaUrl := app.Key("MediaUrl").String()
	file, _ := ctx.FormFile("file")
	bucket := "soulfire-media"

	ext := path.Ext(file.Filename)
	key := utils.Uid("FE") + ext

	dst := "runtime/tmp/imgs/" + key

	err := ctx.SaveUploadedFile(file, dst)
	if err != nil {
		rsp.JsonResonse(ctx, rsp.UploadErr, nil, "")
		return
	}

	img, err := qiniu.Upload(bucket, dst, "feedback/"+key)

	url := MediaUrl + "/" + img

	if err != nil {
		rsp.JsonResonse(ctx, rsp.UploadErr, nil, "")
		return
	}

	_ = os.Remove(dst)

	rsp.JsonResonse(ctx, rsp.OK, url, "")

}

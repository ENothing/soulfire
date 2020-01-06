package user

import (
	"github.com/gin-gonic/gin"
	"soulfire/models"
	"soulfire/pkg/rsp"
	"soulfire/pkg/verify"
	"strconv"
)

type FeedbackForm struct {
	Title   string `json:"title" valid:"Required;" ch:"反馈标题"`
	Content string `json:"content" valid:"Required;" ch:"反馈内容"`
}

func PostFeedback(ctx *gin.Context) {

	userId, _ := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
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
		Pics:    pics,
	}
	err := feedback.Create()

	if err != nil {
		rsp.JsonResonse(ctx, rsp.FeedbackPostFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

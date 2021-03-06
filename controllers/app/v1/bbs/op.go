package bbs

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"os"
	"path"
	"soulfire/models"
	"soulfire/pkg/qiniu"
	"soulfire/pkg/rsp"
	"soulfire/pkg/verify"
	"soulfire/utils"
	"strconv"
)

func Like(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	likes := models.LikeAndUnlike(userId, id, 2)

	if likes == true {

		err := models.ArticleLikeAddOne(id)

		if err != nil {

			rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
			return
		}

		rsp.JsonResonse(ctx, rsp.OK, likes, "")
		return

	}

	err := models.ArticleLikeCutOne(id)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, likes, "")
}

type ArticleForm struct {
	Title   string `json:"title" valid:"Required;" ch:"文章标题"`
	Thumb   string `json:"Thumb" valid:"Required;" ch:"文章封面"`
	Content string `json:"content" valid:"Required;" ch:"文章内容"`
	CateId  int64  `json:"cate_id" valid:"Required;" ch:"文章分类"`
}

func PublishArticle(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)
	thumb := ctx.PostForm("thumb")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	cateId, _ := strconv.ParseInt(ctx.PostForm("cate_id"), 10, 64)
	isPublish, _ := strconv.ParseInt(ctx.PostForm("is_publish"), 10, 64)

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	articleForm := ArticleForm{
		title,
		thumb,
		content,
		cateId,
	}

	message := verify.FormVerify(&articleForm)
	if message != nil {

		rsp.JsonResonse(ctx, rsp.ArticleCreateFailed, nil, message.(string))
		return
	}

	article := models.Article{
		UserId:    userId,
		Thumb:     thumb,
		Title:     title,
		Content:   content,
		CateId:    cateId,
		IsPublish: isPublish,
	}

	err := article.Create()
	if err != nil {

		rsp.JsonResonse(ctx, rsp.ArticleCreateFailed, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func EditArticle(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId,_ := ctx.MustGet("user_id").(int64)
	thumb := ctx.PostForm("thumb")
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	cateId, _ := strconv.ParseInt(ctx.PostForm("cate_id"), 10, 64)
	isPublish, _ := strconv.ParseInt(ctx.PostForm("is_publish"), 10, 64)

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}

	articleForm := ArticleForm{
		title,
		thumb,
		content,
		cateId,
	}

	message := verify.FormVerify(&articleForm)
	if message != nil {

		rsp.JsonResonse(ctx, rsp.ArticleCreateFailed, nil, message.(string))
		return
	}

	_, err := models.GetSelfArticleById(id, userId)

	if err == gorm.ErrRecordNotFound {

		article := models.Article{
			UserId:    userId,
			Thumb:     thumb,
			Title:     title,
			Content:   content,
			CateId:    cateId,
			IsPublish: isPublish,
		}

		err := article.Create()
		if err != nil {

			rsp.JsonResonse(ctx, rsp.ArticleCreateFailed, nil, "")
			return
		}

	}else{
		article := models.Article{
			Title:     title,
			Thumb:     thumb,
			Content:   content,
			CateId:    cateId,
			IsPublish: isPublish,
		}

		err = article.Update(id, userId)
		if err != nil {

			rsp.JsonResonse(ctx, rsp.ArticleUpdateFailed, nil, "")
			return

		}
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func UpdateArticleToPublish(ctx *gin.Context)  {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId,_ := ctx.MustGet("user_id").(int64)
	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	article := models.Article{
		IsPublish: int64(1),
	}

	err := article.Update(id, userId)
	if err != nil {

		rsp.JsonResonse(ctx, rsp.ArticleUpdateFailed, nil, "")
		return

	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func DeleteArticle(ctx *gin.Context) {

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId,_ := ctx.MustGet("user_id").(int64)
	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	article := models.Article{}

	err := article.Delete(id, userId)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.ArticleDeleteFailed, nil, "")
		return

	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func Follow(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)
	followId, _ := strconv.ParseInt(ctx.Param("follow_id"), 10, 64)

	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	if userId == followId {
		rsp.JsonResonse(ctx, rsp.FollowedSelfFailed, nil, "")
		return
	}

	res := models.GetUserFollowById(userId, followId)

	userFollow := models.UserFollow{
		UserId:   userId,
		FollowId: followId,
	}

	if res {

		err := userFollow.Delete()

		if err != nil {

			rsp.JsonResonse(ctx, rsp.FollowCancelFailed, nil, "")
			return

		}
		_ = models.FollowsCutOne(userId)
		_ = models.IsFollowedCutOne(followId)

	} else {

		err := userFollow.Create()
		if err != nil {

			rsp.JsonResonse(ctx, rsp.FollowedFailed, nil, "")
			return

		}
		_ = models.FollowsAddOne(userId)
		_ = models.IsFollowedAddOne(followId)
	}

	rsp.JsonResonse(ctx, rsp.OK, nil, "")

}

func Favor(ctx *gin.Context) {

	userId,_ := ctx.MustGet("user_id").(int64)

	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if userId == int64(0) {
		rsp.JsonResonse(ctx, rsp.PleaseLogin, nil, "")
		return
	}
	favor := models.FavorAndUnFavor(userId, id, 2)

	if favor == true {

		err := models.ArticleFavorAddOne(id)

		if err != nil {

			rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
			return
		}

		rsp.JsonResonse(ctx, rsp.OK, favor, "")
		return

	}

	err := models.ArticleFavorCutOne(id)

	if err != nil {

		rsp.JsonResonse(ctx, rsp.ArticleNotExits, nil, "")
		return
	}

	rsp.JsonResonse(ctx, rsp.OK, favor, "")
}

func Upload(ctx *gin.Context) {

	MediaUrl := viper.GetString("Qiniu.MediaUrl")
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

	img, err := qiniu.Upload(bucket, dst, "bbs/"+key)

	url := MediaUrl + "/" + img

	if err != nil {
		rsp.JsonResonse(ctx, rsp.UploadErr, nil, "")
		return
	}

	_ = os.Remove(dst)

	rsp.JsonResonse(ctx, rsp.OK, url, "")

}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"url-shortener/errorx"
	"url-shortener/log"
	"url-shortener/model"
)

func Redirect(ctx *gin.Context) {
	rsp := response{
		Code: errorx.Success,
		Msg:  errorx.GetMsg(errorx.Success),
		Data: nil,
	}

	shortUrl := ctx.Param("shortUrl")

	url, err := model.FindOneByShortUrl(shortUrl)
	if err != nil || url.IsEmpty() {
		log.Logger.Printf("Short url not found, url: %s", shortUrl)
		rsp.Code = errorx.ShortUrlNotFound
		rsp.Msg = errorx.GetMsg(errorx.ShortUrlNotFound)
		ctx.JSON(http.StatusBadRequest, rsp)
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url.OriginalUrl)
}

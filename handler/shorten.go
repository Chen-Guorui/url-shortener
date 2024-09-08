package handler

import (
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattheath/base62"

	"url-shortener/errorx"
	"url-shortener/log"
	"url-shortener/model"
	"url-shortener/utils"
)

type shortReq struct {
	Url string `json:"url"`
}

func Shorten(ctx *gin.Context) {
	rsp := response{
		Code: errorx.Success,
		Msg:  errorx.GetMsg(errorx.Success),
		Data: nil,
	}

	req := shortReq{}
	err := ctx.BindJSON(&req)
	if err != nil {
		log.Logger.Printf("Failed to parse request, error: %s, param: %+v", err, req)
		rsp.Code = errorx.ParseRequsetError
		rsp.Msg = errorx.GetMsg(errorx.ParseRequsetError)
		ctx.JSON(http.StatusBadRequest, rsp)
		return
	}

	url, err := model.FindOneByOriginalUrl(req.Url)
	if err == nil && !url.IsEmpty() {
		log.Logger.Printf("Original url duplicated, url: %s", req.Url)
		rsp.Code = errorx.OriginalUrlDuplicated
		rsp.Msg = errorx.GetMsg(errorx.OriginalUrlDuplicated)
		ctx.JSON(http.StatusBadRequest, rsp)
		return
	}

	uuid, err := utils.GenerateId()
	if err != nil {
		rsp.Code = errorx.GenerateIdError
		rsp.Msg = errorx.GetMsg(errorx.GenerateIdError)
		ctx.JSON(http.StatusInternalServerError, rsp)
		return
	}
	bigInt := &big.Int{}
	bigInt.SetBytes(uuid[:])
	short := base62.EncodeBigInt(bigInt)
	data := &model.Url{
		Id:          uuid.String(),
		ShortUrl:    short,
		OriginalUrl: req.Url,
	}

	err = model.Insert(data)
	if err != nil {
		log.Logger.Printf("Failed to insert data to database, error: %s, param: %+v", err, data)
		rsp.Code = errorx.InsertToDbError
		rsp.Msg = errorx.GetMsg(errorx.InsertToDbError)
		ctx.JSON(http.StatusInternalServerError, rsp)
		return
	}

	rsp.Data = data
	ctx.JSON(http.StatusOK, rsp)
}

package view

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

const (
	CONTENT_LOCATION = "Content-Location"
)

type SuccessResp struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func MakeSuccessResp(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, SuccessResp{
		Code:   code,
		Status: msg,
		Data:   data,
	})
}

func MakeCreateSuccessResp(c *gin.Context, id string) {
	c.Header(CONTENT_LOCATION, id)
	MakeSuccessResp(c, http.StatusCreated, MsgCreateDataSuccess(c), id)
}

func MakeReadSuccessResp(c *gin.Context, data interface{}) {
	MakeSuccessResp(c, http.StatusOK, MsgGetDataSuccess(c), data)
}

func MakeUpdateSuccessResp(c *gin.Context) {
	MakeSuccessResp(c, http.StatusOK, MsgUpdateDataSuccess(c), reflect.TypeOf(struct{}{}))
}

func MakeDeleteSuccessResp(c *gin.Context) {
	MakeSuccessResp(c, http.StatusOK, MsgDeleteDataSuccess(c), reflect.TypeOf(struct{}{}))
}

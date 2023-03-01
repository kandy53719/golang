package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiCintroller struct {
	BaseController
}

func (ac ApiCintroller) Get(ctx *gin.Context) {
	// ac.Success(ctx, ctx.Request.Method)
	ctx.JSON(http.StatusOK, ctx.Request.Method)
}

func (ac ApiCintroller) Post(ctx *gin.Context) {
	ac.Success(ctx, ctx.Request.Method)
}

func (ac ApiCintroller) Put(ctx *gin.Context) {
	ac.Success(ctx, ctx.Request.Method)
}

func (ac ApiCintroller) Delete(ctx *gin.Context) {
	ac.Success(ctx, ctx.Request.Method)
}

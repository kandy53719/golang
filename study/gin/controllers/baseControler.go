package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

func (bc BaseController) Success(ctx *gin.Context, str string) {
	ctx.String(http.StatusOK, str)
}

func (bc BaseController) Error(ctx *gin.Context, str string) {
	ctx.String(http.StatusOK, str)
}

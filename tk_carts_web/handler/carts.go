package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Numsina/tk_carts/tk_carts_web/domain"
	"github.com/Numsina/tk_carts/tk_carts_web/middlewares"
	"github.com/Numsina/tk_carts/tk_carts_web/service"
	"github.com/Numsina/tk_carts/tk_carts_web/tools"
)

type CartHandler struct {
	svc service.Cart
}

func NewCartHandler(svc service.Cart) *CartHandler {
	return &CartHandler{
		svc: svc,
	}
}

func (h *CartHandler) InitRouter(r *gin.Engine) {
	cartRouter := r.Group("/carts")
	{
		cartRouter.GET("", h.Get)
		cartRouter.POST("", h.Create)
		cartRouter.PUT("", h.Update)
		cartRouter.DELETE("", h.Delete)
		cartRouter.DELETE("/clear", h.Clear)
	}
}

func (c *CartHandler) Create(ctx *gin.Context) {
	var req domain.Carts
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := c.svc.CreateCarts(ctx.Request.Context(), req, claims.UserId)
	if err != nil {
		checkError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
	})
	return
}

func (c *CartHandler) Update(ctx *gin.Context) {
	var req domain.Carts
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := c.svc.UpdateCarts(ctx.Request.Context(), req, claims.UserId)
	if err != nil {
		checkError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
	})
	return
}

func (c *CartHandler) Delete(ctx *gin.Context) {
	type Req struct {
		Id int32 `json:"id"`
	}
	var req Req
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.Result{
			Code: -1,
			Msg:  "参数错误",
		})
		return
	}

	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := c.svc.DeleteCarts(ctx.Request.Context(), req.Id, claims.UserId)
	if err != nil {
		checkError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
	})
	return
}

func (c *CartHandler) Clear(ctx *gin.Context) {
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	err := c.svc.ClearCarts(ctx.Request.Context(), claims.UserId)
	if err != nil {
		checkError(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
	})
	return
}

func (c *CartHandler) Get(ctx *gin.Context) {
	claims := ctx.Value("claims").(*middlewares.UserClaims)
	data, err := c.svc.GetCartsInfo(ctx.Request.Context(), claims.UserId)
	if err != nil {
		checkError(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, tools.Result{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
	return
}

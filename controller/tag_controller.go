package controller

import (
	"gins/data/request"
	"gins/data/response"
	"gins/helper"
	"gins/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TagsController struct {
	tagsService service.TagsService
}

func NewTagasController(service service.TagsService) *TagsController {
	return &TagsController{
		tagsService: service,
	}
}

func (controller *TagsController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagsRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (controller *TagsController) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	updateTagsRequest.Id = id

	controller.tagsService.Update(updateTagsRequest)
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (controller *TagsController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.tagsService.Delete(id)

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (controller *TagsController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	controller.tagsService.FindById(id)

	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    nil,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
func (controller *TagsController) FindAll(ctx *gin.Context) {
	tagsResponse := controller.tagsService.FindAll()
	webResponse := response.WebResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    tagsResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

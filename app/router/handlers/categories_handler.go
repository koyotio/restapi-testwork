package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"rest-api/input"
	"rest-api/response"
	"rest-api/services"
	"strconv"
)

type CategoriesHandlerConfig struct {
	serviceManager *services.ServiceManager
}

func NewCategoriesHandler(serviceManager *services.ServiceManager) *CategoriesHandlerConfig {
	return &CategoriesHandlerConfig{serviceManager: serviceManager}
}

func (h *CategoriesHandlerConfig) Index(ctx *fasthttp.RequestCtx) {
	categories, err := h.serviceManager.CategoryService.GetAllCategories()
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
	} else {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, categories)
	}
}

func (h *CategoriesHandlerConfig) Create(ctx *fasthttp.RequestCtx) {
	var categoryInput input.CategoryInput
	if err := json.Unmarshal(ctx.PostBody(), &categoryInput); err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	if err := categoryInput.Validate(); err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusBadRequest, err.Error())
		return
	}
	category, err := h.serviceManager.CategoryService.CreateCategory(&categoryInput)
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusBadRequest, err.Error())
	} else {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, category)
	}
}

func (h *CategoriesHandlerConfig) Update(ctx *fasthttp.RequestCtx) {
	var categoryInput input.CategoryInput
	id, err := strconv.Atoi(ctx.UserValue(`id`).(string))
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	if err := json.Unmarshal(ctx.PostBody(), &categoryInput); err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	category, err := h.serviceManager.CategoryService.UpdateCategory(id, &categoryInput)
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusNotFound, "category not found")
	} else {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, category)
	}
}

func (h *CategoriesHandlerConfig) View(ctx *fasthttp.RequestCtx) {
	id, err := strconv.Atoi(ctx.UserValue(`id`).(string))
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	category, err := h.serviceManager.CategoryService.GetCategoryById(id)
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusNotFound, err.Error())
	} else {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, category)
	}
}

func (h CategoriesHandlerConfig) Delete(ctx *fasthttp.RequestCtx) {
	id, err := strconv.Atoi(ctx.UserValue(`id`).(string))
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	ok, _ := h.serviceManager.CategoryService.DeleteCategory(id)
	if ok {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, ok)
	} else {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusNotFound, "category not found")
	}
}

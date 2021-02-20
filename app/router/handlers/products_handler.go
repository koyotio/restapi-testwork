package handlers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"rest-api/input"
	"rest-api/response"
	"rest-api/services"
	"strconv"
)

type ProductsHandlerConfig struct {
	serviceManager *services.ServiceManager
}

func NewProductsHandler(serviceManager *services.ServiceManager) *ProductsHandlerConfig {
	return &ProductsHandlerConfig{serviceManager: serviceManager}
}

func (h *ProductsHandlerConfig) Index(ctx *fasthttp.RequestCtx) {
	allProducts, err := h.serviceManager.ProductService.GetAllProducts()
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
	} else {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, allProducts)
	}
}

func (h *ProductsHandlerConfig) Create(ctx *fasthttp.RequestCtx) {
	var productInput input.ProductInput
	if err := json.Unmarshal(ctx.PostBody(), &productInput); err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	if err := productInput.Validate(); err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusBadRequest, err.Error())
		return
	}
	product, err := h.serviceManager.ProductService.CreateProduct(&productInput)
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusBadRequest, err.Error())
	} else {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, product)
	}
}

func (h *ProductsHandlerConfig) Update(ctx *fasthttp.RequestCtx) {
	var productInput input.ProductInput
	id, err := strconv.Atoi(ctx.UserValue(`id`).(string))
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	if err := json.Unmarshal(ctx.PostBody(), &productInput); err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	product, err := h.serviceManager.ProductService.UpdateProduct(id, &productInput)
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusNotFound, "product not found")
	} else {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, product)
	}
}

func (h *ProductsHandlerConfig) View(ctx *fasthttp.RequestCtx) {
	id, err := strconv.Atoi(ctx.UserValue(`id`).(string))
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	product, err := h.serviceManager.ProductService.GetProductById(id)
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusNotFound, err.Error())
	} else {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, product)
	}
}

func (h *ProductsHandlerConfig) Delete(ctx *fasthttp.RequestCtx) {
	id, err := strconv.Atoi(ctx.UserValue(`id`).(string))
	if err != nil {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusInternalServerError, err.Error())
		return
	}
	ok, _ := h.serviceManager.ProductService.DeleteProduct(id)
	if ok {
		response.NewJSONSuccessResponse(ctx, fasthttp.StatusOK, ok)
	} else {
		response.NewJSONErrorResponse(ctx, fasthttp.StatusNotFound, "product not found")
	}
}

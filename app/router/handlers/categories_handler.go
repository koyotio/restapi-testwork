package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"rest-api/services"
)

type CategoriesHandlerConfig struct {
	serviceManager *services.ServiceManager
}

func NewCategoriesHandler(serviceManager *services.ServiceManager) *CategoriesHandlerConfig {
	return &CategoriesHandlerConfig{serviceManager: serviceManager}
}

func (h CategoriesHandlerConfig) Index(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(200)
	ctx.SetContentType(`application/json`)
	categories, err := h.serviceManager.CategoryService.GetAllCategories()
	if err != nil {
		response, err := json.Marshal(err)
		if err != nil {
			ctx.WriteString(fmt.Sprintf(`some error occurred when trying to send response in JSON: %s`, err.Error()))
			return
		}
		ctx.Write(response)
		return
	}
	categoriesJSON, err := json.Marshal(categories)
	if err != nil {
		response, err := json.Marshal(err)
		if err != nil {
			ctx.WriteString(fmt.Sprintf(`some error occurred when trying to send response in JSON: %s`, err.Error()))
			return
		}
		ctx.Write(response)
		return
	}
	ctx.Write(categoriesJSON)
}

func (h CategoriesHandlerConfig) Create(ctx *fasthttp.RequestCtx) {

}

func (h CategoriesHandlerConfig) Update(ctx *fasthttp.RequestCtx) {

}

func (h CategoriesHandlerConfig) View(ctx *fasthttp.RequestCtx) {

}

func (h CategoriesHandlerConfig) Delete(ctx *fasthttp.RequestCtx) {

}

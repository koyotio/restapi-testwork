package handlers

import (
	"github.com/valyala/fasthttp"
	"rest-api/services"
)

type ProductsHandlerConfig struct {
	serviceManager *services.ServiceManager
}

func NewProductsHandler(serviceManager *services.ServiceManager) *ProductsHandlerConfig {
	return &ProductsHandlerConfig{serviceManager: serviceManager}
}

func (h ProductsHandlerConfig) Index(ctx *fasthttp.RequestCtx) {

}

func (h ProductsHandlerConfig) Create(ctx *fasthttp.RequestCtx) {

}

func (h ProductsHandlerConfig) Update(ctx *fasthttp.RequestCtx) {

}

func (h ProductsHandlerConfig) View(ctx *fasthttp.RequestCtx) {

}

func (h ProductsHandlerConfig) Delete(ctx *fasthttp.RequestCtx) {

}

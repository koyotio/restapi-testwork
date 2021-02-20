package router

import (
	fasthttpRouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"rest-api/router/handlers"
	"rest-api/services"
)

type Config struct {
	serviceManager *services.ServiceManager
	HTTPRouter     *fasthttpRouter.Router
}

func NewRouter(serviceManager *services.ServiceManager) *Config {
	return &Config{
		serviceManager: serviceManager,
		HTTPRouter:     fasthttpRouter.New(),
	}
}

func (r Config) InitRoutes() {
	ch := handlers.NewCategoriesHandler(r.serviceManager)
	ph := handlers.NewProductsHandler(r.serviceManager)

	r.HTTPRouter.POST("/categories", ch.Create)
	r.HTTPRouter.GET("/categories", ch.Index)
	r.HTTPRouter.GET("/categories/{id}", ch.View)
	r.HTTPRouter.PUT("/categories/{id}", ch.Update)
	r.HTTPRouter.DELETE("/categories/{id}", ch.Delete)
	r.HTTPRouter.POST("/products", ph.Create)
	r.HTTPRouter.GET("/products", ph.Index)
	r.HTTPRouter.GET("/products/{id}", ph.View)
	r.HTTPRouter.PUT("/products/{id}", ph.Update)
	r.HTTPRouter.DELETE("/products/{id}", ph.Delete)
}

func (r Config) GetHandler() func(ctx *fasthttp.RequestCtx) {
	r.InitRoutes()
	return r.HTTPRouter.Handler
}

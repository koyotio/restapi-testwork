package response

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

const (
	StatusOK    string = "OK"
	StatusError string = "ERROR"
)

type JSONResponseData struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
}

func NewJSONSuccessResponse(ctx *fasthttp.RequestCtx, code int, data interface{}) {
	ctx.Response.Header.SetStatusCode(code)
	ctx.Response.Header.SetContentType("application/json")
	response := JSONResponseData{
		Status: StatusOK,
		Data:   data,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		logrus.Errorf("error while preparing json response: %s", err.Error())
		ctx.Error("some error occurred when preparing response", fasthttp.StatusInternalServerError)
	}
}

func NewJSONErrorResponse(ctx *fasthttp.RequestCtx, code int, msg string) {
	logrus.Error(msg)
	ctx.Response.Header.SetStatusCode(code)
	ctx.Response.Header.SetContentType("application/json")
	response := JSONResponseData{
		Status: StatusError,
		Error:  msg,
	}
	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		logrus.Errorf("some error occurred when preparing response: %s", err.Error())
		ctx.Error("some error occurred when preparing response", fasthttp.StatusInternalServerError)
	}
}

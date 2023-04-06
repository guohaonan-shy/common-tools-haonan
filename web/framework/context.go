package framework

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Response   http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int64
}

func NewContext(resp http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Response: resp,
		Request:  req,
		Path:     req.URL.Path,
		Method:   req.Method,
	}
}

func (ctx *Context) Query(key string) string {
	// parse parameters in url, such as GET
	return ctx.Request.URL.Query().Get(key)
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Request.FormValue(key)
}

func (ctx *Context) SetStatusCode(code int64) {
	ctx.StatusCode = code
	ctx.Response.WriteHeader(int(code))
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.Response.Header().Set(key, value)
}

func (ctx *Context) String(code int64, format string, values ...any) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.SetStatusCode(code)
	_, _ = ctx.Response.Write([]byte(fmt.Sprintf(format, values...)))
}

func (ctx *Context) JSON(code int64, obj any) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.SetStatusCode(code)
	encoder := json.NewEncoder(ctx.Response)
	_ = encoder.Encode(obj)
}

func (ctx *Context) Data(code int64, data []byte) {
	ctx.SetStatusCode(code)
	_, _ = ctx.Response.Write(data)
}

func (ctx *Context) HTML(code int64, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.SetStatusCode(code)
	_, _ = ctx.Response.Write([]byte(html))
}

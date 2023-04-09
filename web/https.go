package main

import (
	"fmt"
	"github.com/common-tools-haonan/web/framework"
	"strings"
)

//type HTTPEngine struct {
//}
//
//func NewHttpEngine() *HTTPEngine {
//	return &HTTPEngine{}
//}

func main() {

	//http.HandleFunc("/v1", V1Handler)
	//http.HandleFunc("/v2", V2Handler)
	//http.HandleFunc("/v3", V3Hanlder)

	//engine := NewHttpEngine()
	//log.Fatal(http.ListenAndServe(":9999", engine))

	engine := framework.NewShyHTTP()
	engine.POST("/v1", V1Handler)
	//engine.GET("/testing/dynamic/:username", V2Handler)
	engine.GET("/testing/dynamic/*", V2Handler)

	engine.Run(":9999")

}

//func (engine *HTTPEngine) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
//
//	switch req.URL.Path {
//	case "/v1":
//		V1Handler(resp, req)
//	case "/v2":
//		V2Handler(resp, req)
//	case "/v3":
//		V3Handler(resp, req)
//	default:
//		fmt.Fprintf(resp, "path hasn't registered\n")
//	}
//
//}

//func V1Handler(resp http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(resp, "[V1Handler] exec success\n")
//}
//
//func V2Handler(resp http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(resp, "URL.PATH=%q\n", req.URL.Path)
//}
//
//func V3Handler(resp http.ResponseWriter, req *http.Request) {
//	for k, v := range req.Header {
//		fmt.Fprintf(resp, "[Header] key=%q, value=%q\n", k, v)
//	}
//}

func V1Handler(ctx *framework.Context) {
	//fmt.Fprintf(ctx.Response, "[V1Handler] exec success\n")
	uid := ctx.PostForm("uid")
	password := ctx.PostForm("password")
	ctx.JSON(200, map[string]any{
		"uid":      uid,
		"password": password,
	})
}

func V2Handler(ctx *framework.Context) {
	//fmt.Fprintf(ctx.Response, "URL.PATH=%q\n", ctx.Request.URL.Path)
	ctx.String(200, "URL.PATH=%q\n", ctx.Request.URL.Path)
}

func V3Handler(ctx *framework.Context) {
	//for k, v := range ctx.Request.Header {
	//	fmt.Fprintf(ctx.Response, "[Header] key=%q, value=%q\n", k, v)
	//}
	builder := strings.Builder{}
	builder.WriteString("<h1> ")
	builder.WriteString(fmt.Sprintf("user:%s ", ctx.Query("username")))
	builder.WriteString(fmt.Sprintf("password:%s ", ctx.Query("password")))
	builder.WriteString("<h1> \n")
	ctx.HTML(200, builder.String())
}

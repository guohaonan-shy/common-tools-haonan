package main

import (
	"fmt"
	"github.com/common-tools-haonan/web/framework"
	"net/http"
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
	engine.GET("/v1", V1Handler)
	engine.GET("/v2", V2Handler)
	engine.POST("/v3", V3Handler)

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

func V1Handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "[V1Handler] exec success\n")
}

func V2Handler(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "URL.PATH=%q\n", req.URL.Path)
}

func V3Handler(resp http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(resp, "[Header] key=%q, value=%q\n", k, v)
	}
}

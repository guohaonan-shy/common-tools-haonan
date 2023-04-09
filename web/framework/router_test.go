package framework

import (
	"io"
	"net/http"
	"testing"
)

//var Engine *ShyHTTP

//func init() {
//	Engine = NewShyHTTP()
//	testHandler := func(ctx *Context) {
//		fmt.Fprintf(ctx.Response, "[V1Handler] exec success\n")
//		//uid := ctx.PostForm("uid")
//		//password := ctx.PostForm("password")
//		//ctx.JSON(200, map[string]any{
//		//	"uid":      uid,
//		//	"password": password,
//		//})
//	}
//	Engine.GET("/testing/v1/:username", testHandler)
//	Engine.Run(":9999")
//}

func Test_Router(t *testing.T) {
	router := NewRouter()

	router.AddRoute("POST", "/testing/v1", nil)
	router.AddRoute("GET", "/testing/v2", nil)
	router.AddRoute("GET", "/testing/v3", nil)

	fn := func(ctx *Context) {
		//fmt.Fprintf(ctx.Response, "URL.PATH=%q\n", ctx.Request.URL.Path)
		ctx.String(200, "URL.PATH=%q\n", ctx.Request.URL.Path)
	}

	router.AddRoute("GET", "/testing/dynamic/:username", fn)

	router.RouteFunc("GET", "/testing/dynamic/ghn980421")

	t.Logf("router:%+v", router)
}

func Test_CurlDynamicRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:9999/testing/v1/ghn980421", nil)
	if err != nil {
		t.Errorf("[http.NewRequest] exec failed, err:%s", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("[http] url:%s exec failed, err:%s", req.URL, err)
		return
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == http.StatusOK {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("[io] Read response body failed, err:%s", err)
			return
		} else {
			t.Log(bytes)
		}
	} else {
		t.Errorf("[https] exec failed, resp:%+v", resp)
	}

}

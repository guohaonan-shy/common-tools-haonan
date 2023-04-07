package framework

import (
	"testing"
)

func Test_Router(t *testing.T) {
	router := NewRouter()

	router.AddRoute("POST", "/testing/v1", nil)
	router.AddRoute("GET", "/testing/v2", nil)
	router.AddRoute("GET", "/testing/v3", nil)

	t.Logf("router:%+v", router)
}

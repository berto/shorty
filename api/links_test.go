package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListLinks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testRouter := SetupRouter()

	req, err := http.NewRequest("GET", "/api/v1/links", nil)
	if err != nil {
		t.Errorf("Get failed with error %d.", err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)

	if resp.Code != 200 {
		t.Errorf("/api/v1/links failed with error code %d.", resp.Code)
	}
}

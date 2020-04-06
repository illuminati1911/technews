package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	deliveryHttp "github.com/illuminati1911/technews/service-auth/auth/delivery/http"
	"github.com/illuminati1911/technews/service-auth/auth/repository/mock"
	"github.com/illuminati1911/technews/service-auth/auth/service"
	"gopkg.in/go-playground/assert.v1"
)

func getAuthRouter() *gin.Engine {
	router := gin.Default()
	repo := mock.NewMockAuthRepository()
	serv := service.NewAuthService(repo)
	deliveryHttp.NewAuthHTTPHandler(serv, router)
	return router
}

func getResponse(t *testing.T, router *gin.Engine, method string, route string, payload string) *httptest.ResponseRecorder {
	jsonReq := []byte(payload)
	req, err := http.NewRequest(method, route, bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Errorf("Could not create a request")
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	return res
}

func TestCreateUser(t *testing.T) {
	router := getAuthRouter()
	res := getResponse(t, router, "POST", "/create", `{"username": "testuser", "password": "testpassword"}`)
	assert.Equal(t, res.Code, 200)
}

func TestCreateUserFailure(t *testing.T) {
	router := getAuthRouter()
	res := getResponse(t, router, "POST", "/create", `{"username": "testuser"}`)
	assert.Equal(t, res.Code, 400)
}

func TestCreateUserAndLogin(t *testing.T) {
	router := getAuthRouter()
	res := getResponse(t, router, "POST", "/create", `{"username": "testuser", "password": "testpassword"}`)
	assert.Equal(t, res.Code, 200)
	res = getResponse(t, router, "POST", "/login", `{"username": "testuser", "password": "testpassword"}`)
	assert.Equal(t, res.Code, 200)
	type tokenResp struct {
		Token *string `json:"token"`
	}
	var token tokenResp
	if err := json.Unmarshal(res.Body.Bytes(), &token); err != nil {
		t.Error("Could not parse JSON")
	}
	if token.Token == nil {
		t.Error("Did not receive token")
	}
}

func TestCreateUserAndLoginWithWrongPassword(t *testing.T) {
	router := getAuthRouter()
	res := getResponse(t, router, "POST", "/create", `{"username": "testuser", "password": "testpassword"}`)
	assert.Equal(t, res.Code, 200)
	res = getResponse(t, router, "POST", "/login", `{"username": "testuser", "password": "testpassword123"}`)
	assert.Equal(t, res.Code, 401)
	type errorResp struct {
		Error *string `json:"error"`
	}
	var errRes errorResp
	if err := json.Unmarshal(res.Body.Bytes(), &errRes); err != nil {
		t.Error("Could not parse JSON")
	}
	if errRes.Error == nil {
		t.Error("Expected to contain error message")
	}
}

func TestLoginWithoutPassword(t *testing.T) {
	router := getAuthRouter()
	res := getResponse(t, router, "POST", "/login", `{"username": "testuser"}`)
	assert.Equal(t, res.Code, 400)
}

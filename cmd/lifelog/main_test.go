package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/Viovie-com/lifelog/internal/db"
	"github.com/Viovie-com/lifelog/internal/server"
)

func performApiRequest(r http.Handler, method, path string, body *gin.H) *httptest.ResponseRecorder {
	var jsonValue []byte
	if body != nil {
		jsonValue, _ = json.Marshal(body)
	}
	req, _ := http.NewRequest(method, "/api/v1"+path, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestApiMemberRegister(t *testing.T) {
	sqlDb, _, _ := sqlmock.New()
	db.SetMockDb(sqlDb)

	router := server.SetupRouter()
	body := gin.H{
		"name":     "Test Name",
		"email":    "test@123.c",
		"password": "123",
	}
	w := performApiRequest(router, http.MethodPost, "/member/", &body)

	assert.Equal(t, http.StatusOK, w.Code)
}

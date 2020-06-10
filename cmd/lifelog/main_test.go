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

func performApiRequest(r http.Handler, method, path string, headers *map[string]string, body *gin.H) *httptest.ResponseRecorder {
	var jsonValue []byte
	if body != nil {
		jsonValue, _ = json.Marshal(body)
	}
	req, _ := http.NewRequest(method, "/api/v1"+path, bytes.NewBuffer(jsonValue))
	if headers != nil {
		for key, val := range *headers {
			req.Header.Set(key, val)
		}
	}

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
	w := performApiRequest(router, http.MethodPost, "/member/", nil, &body)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiMemberUpdate(t *testing.T) {
	sqlDb, mockDb, _ := sqlmock.New()
	mockDb.ExpectQuery("^SELECT (.*)").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockDb.ExpectQuery("^SELECT (.*)").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockDb.ExpectClose()
	db.SetMockDb(sqlDb)

	router := server.SetupRouter()
	headers := map[string]string{
		"Authorization": "bearer 123",
	}
	body := gin.H{
		"name": "Test Name",
	}
	w := performApiRequest(router, http.MethodPut, "/member/", &headers, &body)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiAuthLogin(t *testing.T) {
	sqlDb, mockDb, _ := sqlmock.New()
	mockDb.ExpectQuery("^SELECT (.*)").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	db.SetMockDb(sqlDb)

	router := server.SetupRouter()
	body := gin.H{
		"email":    "test@123.c",
		"password": "123",
	}
	w := performApiRequest(router, http.MethodPost, "/auth/", nil, &body)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiPostAdd(t *testing.T) {
	sqlDb, mockDb, _ := sqlmock.New()
	mockDb.ExpectQuery("^SELECT (.*)").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockDb.ExpectQuery("^SELECT (.*)").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockDb.ExpectClose()
	db.SetMockDb(sqlDb)

	router := server.SetupRouter()
	headers := map[string]string{
		"Authorization": "bearer 123",
	}
	body := gin.H{
		"title":      "title",
		"content":    "test content",
		"categoryId": 1,
		"draft":      true,
		"tags":       []string{"Test", "123"},
	}
	w := performApiRequest(router, http.MethodPost, "/post/", &headers, &body)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiPostList(t *testing.T) {
	sqlDb, _, _ := sqlmock.New()
	db.SetMockDb(sqlDb)

	router := server.SetupRouter()
	w := performApiRequest(router, http.MethodGet, "/post/", nil, nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

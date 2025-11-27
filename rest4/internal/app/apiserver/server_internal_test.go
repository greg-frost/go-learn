package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-learn/rest4/internal/app/store/teststore"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := NewServer(teststore.New())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)

	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Equal(t, rec.Body.String(), "Привет, пользователи!")
}

package controllers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooksControllers(t *testing.T) {
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, GetBooksControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "success", result["status"])
}

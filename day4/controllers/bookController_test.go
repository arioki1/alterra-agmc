package controllers

import (
	"encoding/json"
	"github.com/arioki1/alterra-agmc/day4/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetBooksControllers(t *testing.T) {
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, GetBooksControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "success", result["status"])
}

func TestCreateBookControllersInvalidJson(t *testing.T) {
	//setup echo context
	e := echo.New()

	//create json body
	body := `{"title":1","isbn":"test","writer":"test"}`

	//setup request
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, CreateBookControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "Invalid JSON body", result["status"])
}

func TestCreateBookControllersBadRequest(t *testing.T) {
	//setup echo context
	e := echo.New()

	body := models.Book{
		Title: "test",
		Isbn:  "1234567890",
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, CreateBookControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "writer is required", result["status"])
}

func TestCreateBookControllersSuccess(t *testing.T) {
	//setup echo context
	e := echo.New()

	body := models.Book{
		Title:  "test",
		Isbn:   "1234567890",
		Writer: "test",
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, CreateBookControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "created", result["status"])
}

func TestGetBookByIdControllersInvalidParams(t *testing.T) {
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, GetBookByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "Invalid params", result["status"])
}

func TestGetBookByIdControllersSuccess(t *testing.T) {
	// add books
	book := models.Book{
		Id:     1,
		Title:  "test",
		Isbn:   "1234567890",
		Writer: "test",
	}

	books = append(books, book)
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, GetBookByIdControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "success", result["status"])
}

func TestGetBookByIdControllersNotFound(t *testing.T) {
	// add books
	book := models.Book{
		Id:     1,
		Title:  "test",
		Isbn:   "1234567890",
		Writer: "test",
	}

	books = append(books, book)
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("20")

	//test
	assert.NoError(t, GetBookByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "book not found", result["status"])
}

func TestUpdateBookByIdControllersInvalidParams(t *testing.T) {
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodPut, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, UpdateBookByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "Invalid params", result["status"])
}

func TestUpdateBookByIdControllersInvalidJson(t *testing.T) {
	//setup echo context
	e := echo.New()

	//create json body
	body := `{"title":1","isbn":"test","writer":"test"}`

	//setup request
	req := httptest.NewRequest(http.MethodPut, "/books", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, UpdateBookByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "Invalid JSON body", result["status"])
}

func TestUpdateBookByIdControllersSuccess(t *testing.T) {
	//setup echo context
	e := echo.New()

	body := models.Book{
		Id:     1,
		Title:  "test",
		Isbn:   "1234567890",
		Writer: "test",
	}
	//add book
	books = append(books, body)

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/books", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, UpdateBookByIdControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "updated", result["status"])
}

func TestUpdateBookByIdControllersNotFound(t *testing.T) {
	//setup echo context
	e := echo.New()

	body := models.Book{
		Id:     1,
		Title:  "test",
		Isbn:   "1234567890",
		Writer: "test",
	}
	//add book
	books = append(books, body)

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/books", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("20")

	//test
	assert.NoError(t, UpdateBookByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "book not found", result["status"])
}

func TestDeleteBookByIdControllersInvalidParams(t *testing.T) {
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, DeleteBookByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "Invalid params", result["status"])
}

func TestDeleteBookByIdControllersSuccess(t *testing.T) {
	//setup echo context
	e := echo.New()

	body := models.Book{
		Id:     1,
		Title:  "test",
		Isbn:   "1234567890",
		Writer: "test",
	}
	//add book
	books = append(books, body)

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, DeleteBookByIdControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "deleted", result["status"])
}

func TestDeleteBookByIdControllersNotFound(t *testing.T) {
	//setup echo context
	e := echo.New()

	body := models.Book{
		Id:     1,
		Title:  "test",
		Isbn:   "1234567890",
		Writer: "test",
	}
	//add book
	books = append(books, body)

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues("20")

	//test
	assert.NoError(t, DeleteBookByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "book not found", result["status"])
}

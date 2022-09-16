package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/arioki1/alterra-agmc/day4/config"
	"github.com/arioki1/alterra-agmc/day4/lib/database/seeder"
	"github.com/arioki1/alterra-agmc/day4/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// init function testing
func setupTest(t *testing.T) {
	//load env
	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error3 loading .env file")
	}

	//setup database
	config.InitDB()

	// clear database
	s := seeder.NewUserSeeder()
	fmt.Println(s)
	s.Delete()
	s.Seed()
}

func TestGetUserControllers(t *testing.T) {
	setupTest(t)
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, GetUserControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "success", result["status"])
}

func TestCreateUserControllersSuccess(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Users{
		Name:     "yoga",
		Email:    "yoga@mail.com",
		Password: "pwd",
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, CreateUserControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "created", result["status"])
}

func TestCreateUserControllersFailedWhenUserNotInputEmail(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Users{
		Name:     "yoga",
		Email:    "",
		Password: "pwd",
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, CreateUserControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "email is required", result["status"])

}

func TestGetUserByIdControllersSuccess(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("1")

	//test
	assert.NoError(t, GetUserByIdControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "success", result["status"])
}

func TestGetUserByIdControllersNotFound(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("10")

	//test
	assert.NoError(t, GetUserByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "user not found", result["status"])
}

func TestUpdateUserByIdControllersSuccess(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Users{
		Name:     "yoga",
		Email:    "yoga@mail.com",
		Password: "pwd",
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("1")

	//set user id
	c.Set("userId", 1)

	//test
	assert.NoError(t, UpdateUserByIdControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "success", result["status"])
}

func TestUpdateUserByIdControllersNotOwnedID(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Users{
		Name:     "yoga",
		Email:    "yoga@mail.com",
		Password: "pwd",
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("1")

	//set user id
	c.Set("userId", 2)

	//test
	assert.NoError(t, UpdateUserByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "you can edit your own data only", result["status"])
}

func TestUpdateUserByIdControllersNotFound(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Users{
		Name:     "yoga",
		Email:    "yoga@mail.com",
		Password: "pwd",
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("10")

	//set user id
	c.Set("userId", 10)

	//test
	assert.NoError(t, UpdateUserByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, fmt.Sprintf("row with id=%v  cannot be update because it doesn't exist", 10), result["status"])
}

package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/arioki1/alterra-agmc/day7/config"
	"github.com/arioki1/alterra-agmc/day7/lib/database/seeder"
	"github.com/arioki1/alterra-agmc/day7/models"
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

func TestLoginUsersControllersSuccess(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Users{
		Email:    "test1@mail.com",
		Password: "test1",
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	assert.NoError(t, LoginUsersControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	dataUsers := result["users"].(map[string]interface{})
	assert.Equal(t, "success login", result["status"])
	assert.NotNil(t, dataUsers["name"])
	assert.NotEmpty(t, dataUsers["name"])
	assert.Equal(t, body.Email, dataUsers["email"])
	assert.Nil(t, dataUsers["password"])
	assert.NotNil(t, dataUsers["token"])
	assert.NotEmpty(t, dataUsers["token"])
}

func TestLoginUsersControllersWrongEmailOrPassword(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//create json body
	body := models.Users{
		Email:    "test1@mail.com",
		Password: "test2",
	}

	//setup request
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(b)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	err := LoginUsersControllers(c)
	he, ok := err.(*echo.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, http.StatusBadRequest, he.Code)
	assert.Equal(t, "wrong email or password", he.Message)
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

func TestGetUserControllersFailedDBNotConnect(t *testing.T) {
	setupTest(t)
	db, err := config.DB.DB()
	assert.NoError(t, err)
	assert.NoError(t, db.Close())
	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//test
	err = GetUserControllers(c)
	assert.Error(t, err)
	he, ok := err.(*echo.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, http.StatusInternalServerError, he.Code)
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

func TestDeleteUserByIdControllersSuccess(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("1")

	//set user id
	c.Set("userId", 1)

	//test
	assert.NoError(t, DeleteUserByIdControllers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "deleted", result["status"])
}

func TestDeleteUserByIdControllersNotOwnedID(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("1")

	//set user id
	c.Set("userId", 2)

	//test
	assert.NoError(t, DeleteUserByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, "you can delete your own data only", result["status"])
}

func TestDeleteUserByIdControllersNotFound(t *testing.T) {
	setupTest(t)

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set params
	c.SetParamNames("id")
	c.SetParamValues("10")

	//set user id
	c.Set("userId", 10)

	//test
	assert.NoError(t, DeleteUserByIdControllers(c))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
	result := map[string]interface{}{}
	assert.NoError(t, json.NewDecoder(rec.Body).Decode(&result))
	assert.Equal(t, fmt.Sprintf("row with id=%v  cannot be delete because it doesn't exist", 10), result["status"])
}

func TestGetUserId(t *testing.T) {

	//setup echo context
	e := echo.New()

	//setup request
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	//set user id
	c.Set("userId", 1)

	//test
	assert.Equal(t, 1, getUserId(c))
}

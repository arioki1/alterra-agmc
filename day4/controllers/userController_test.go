package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/arioki1/alterra-agmc/day4/lib/database/seeder"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

// init function testing
func TestInit(t *testing.T) {
	//load env
	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error3 loading .env file")
	}

	// setup database
	s := seeder.NewUserSeeder()
	fmt.Println(s)
	s.Delete()
	s.Seed()
}

func TestGetUserControllers(t *testing.T) {
	//setup echo context
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetBooksControllers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert  is body json
		result := map[string]interface{}{}
		if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
			log.Fatalln(err)
		}
		assert.Equal(t, "success", result["status"])
	}
}

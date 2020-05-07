package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T){
	//setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)


	// Assertions
	if assert.NoError(t, Hello()(c)) {
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, "heelo world!", rec.Body.String())
	}
}

func TestHello2(t *testing.T){
	//setup
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)


	// Assertions
	if assert.NoError(t, Hello2()(c)) {
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, "heelo2 world!", rec.Body.String())
	}
}

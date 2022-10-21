package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

// ContextAndRecorderGet context and recorder for Get method
func ContextAndRecorderGet(path string, e *echo.Echo) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	return ctx, rec
}

// ContextAndRecorderPut context and recorder Put method
func ContextAndRecorderPut(path, file string, e *echo.Echo) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPut, path, strings.NewReader(RemoveEnterAndTabs(Read(file))))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	return ctx, rec
}

// ContextAndRecorderPost context and recorder Put method
func ContextAndRecorderPost(path, file string, e *echo.Echo) (echo.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(RemoveEnterAndTabs(Read(file))))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	return ctx, rec
}

// Read reads json data from file
func Read(file string) string {
	content, err := os.Open(file)
	if err != nil {
		return ""
	}
	defer content.Close()
	byteValue, err := ioutil.ReadAll(content)
	if err != nil {
		return ""
	}
	return RemoveEnterAndTabs((string(byteValue)))
}

// RemoveEnterAndTabs removes enters and tabs
func RemoveEnterAndTabs(s string) string {
	return fmt.Sprintf("%s\n", strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), "\t", ""), " ", ""))
}

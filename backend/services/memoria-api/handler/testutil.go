package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func newGinContext(method string, path string) (c *gin.Context) {
	w := httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	req, err := http.NewRequest(
		method,
		path,
		nil,
	)
	if err != nil {
		panic(err)
	}

	c.Request = req
	return
}

func newGinContextWithBody(method string, path string, body any) (c *gin.Context) {
	w := httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	c.Request, err = http.NewRequest(
		method,
		path,
		bytes.NewReader(jsonBytes),
	)
	if err != nil {
		panic(err)
	}

	return
}

func verifyJSONEncoding[T any](data any, decoded *T) {
	resJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(resJson, &decoded)
	if err != nil {
		panic(err)
	}
}

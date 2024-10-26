package data

import (
	"encoding/json"
	"github.com/steinfletcher/apitest"
	"io"
)

type HttpHeaderOption func(map[string]string)

func HttpHeaders(options ...HttpHeaderOption) map[string]string {
	headers := map[string]string{}
	for _, option := range options {
		option(headers)
	}
	return headers
}

func BearerToken(token string) HttpHeaderOption {
	return func(headers map[string]string) {
		headers["Authentication"] = "Bearer " + token
	}
}

type Request[T any] struct {
	Request T
}

func (r Request[T]) Body() string {
	body, _ := json.Marshal(r.Request)
	return string(body)
}

type Response[T any] struct {
	Result apitest.Result
}

func (r Response[T]) StatusCode() int {
	return r.Result.Response.StatusCode
}

func (r Response[T]) Body() (t T) {
	body, _ := io.ReadAll(r.Result.Response.Body)
	_ = json.Unmarshal(body, &t)
	return
}

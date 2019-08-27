package context

import (
	"context"
	"encoding/json"
	"net/http"
)

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	ctx      context.Context
}

func New(ctx context.Context, w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ctx:      ctx,
		response: w,
		request:  r,
	}
}

func (c *Context) SetRW(w http.ResponseWriter, r *http.Request) {
	c.response = w
	c.request = r
}

func (c *Context) Response() http.ResponseWriter {
	return c.response
}

func (c *Context) Request() *http.Request {
	return c.request
}

func (c *Context) Context() context.Context {
	return c.ctx
}

func (c *Context) Json(data interface{}) {
	c.response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		c.response.WriteHeader(500)
		c.response.Write([]byte(err.Error()))
		return
	}

	c.response.Write(jsonBytes)
}

func (c *Context) ParamsGetAll() (values map[string]string) {
	params := c.request.URL.Query()

	values = map[string]string{}
	for key, item := range params {
		values[key] = item[0]
	}
	return values
}
func (c *Context) ParamsGet(key string) (value string) {
	params := c.ParamsGetAll()

	value, ok := params[key]

	if ok == true {
		return value
	}
	return ""
}

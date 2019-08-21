package context

import (
	"context"
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

package api

import (
    "github.com/kataras/iris"
)

type StatusAPI struct {
	  *iris.Context
}

// GET /status
func (s StatusAPI) Get() {
    m := formJson("getStatus", "", s.RequestCtx.Request.Body())
    s.Write(reqCore(m))
}


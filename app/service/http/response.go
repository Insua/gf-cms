package http

import (
	"net/http"

	"github.com/gogf/gf/util/gvalid"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type PaginateMeta struct {
	CurrentPage uint `json:"current_page"`
	PerPage     uint `json:"per_page"`
	Total       uint `json:"total"`
}

func Error422(r *ghttp.Request, errors interface{}) {
	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatusExit(http.StatusUnprocessableEntity, g.Map{"errors": errors})
}

func Error500(r *ghttp.Request) {
	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatusExit(http.StatusInternalServerError, g.Map{"errors": "server error"})
}

func ErrorValid(r *ghttp.Request, err *gvalid.Error) {
	if err != nil {
		errors := make([]string, 0)
		for _, v := range err.Maps() {
			for _, v := range v {
				errors = append(errors, v)
			}
		}
		Error422(r, errors)
	}
}

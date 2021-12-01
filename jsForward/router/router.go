package router

import (
	"jsForward/app/api"
	"jsForward/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(service.MiddlewareCORS)
		group.POST("/request",api.Home.RequestData)
		group.POST("/response",api.Home.ResponseData)
	})
}

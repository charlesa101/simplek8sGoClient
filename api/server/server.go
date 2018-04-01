package server

import (
	"github.com/k8sApi/api/core"
	"github.com/ericchiang/k8s"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/k8sApi/api/route"
)

type HttpServer struct {
	*core.KubernetesContext
}

func HttpHandler(context context.Context, k8sClient *k8s.Client) *HttpServer {
	return &HttpServer{
		KubernetesContext: &core.KubernetesContext{
			Context: context,
			Client:  k8sClient,
		},
	}
}


func (s *HttpServer) Bind(addr ...string) error {

	r := gin.Default()
	r.Use(core.HttpHeaderWithCorsSupport())
	{
		//Register API routes
		route.Routes(r.Group("/api/"), s.KubernetesContext)
	}

	return r.Run(addr...)
}
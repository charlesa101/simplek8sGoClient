package route

import (
	"github.com/k8sApi/api/core"
	"github.com/gin-gonic/gin"
)

type Service struct {
	ctx *core.KubernetesContext
}

func Routes (r *gin.RouterGroup, ctx *core.KubernetesContext){

	s := &Service{ ctx: ctx, }

	r.GET("/nodes", s.nodes )
	r.GET("/nodes/count", s.nodeCount )
	r.GET("/pods/:namespace", s.pods )
	r.GET("/pods/:namespace/count", s.podCount )


}
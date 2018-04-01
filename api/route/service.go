package route

import (
	"net/http"
	"github.com/gin-gonic/gin"
	corev1 "github.com/ericchiang/k8s/apis/core/v1"
	"fmt"
)

func (s *Service) nodes( c *gin.Context) {
	var nodes corev1.NodeList
	namespace := ""
	s.ctx.Client.List(c, namespace, &nodes)
	c.JSON(http.StatusOK, nodes.Items)
}

func (s *Service) nodeCount( c *gin.Context) {
	var nodes corev1.NodeList
	namespace := ""
	s.ctx.Client.List(c, namespace, &nodes)
	c.JSON(http.StatusOK, fmt.Sprintf("{ 'nodes' : %d }", len(nodes.Items)))
}

func (s *Service) pods( c *gin.Context) {
	var pods corev1.PodList
	namespace := c.Param("namespace")
	s.ctx.Client.List(c, namespace, &pods)
	c.JSON(http.StatusOK, pods.Items)
}


func (s *Service) podCount( c *gin.Context) {
	var pods corev1.PodList
	namespace := c.Param("namespace")
	s.ctx.Client.List(c, namespace, &pods)
	c.JSON(http.StatusOK, fmt.Sprintf("{ 'pods' : %d }", len(pods.Items)))
}
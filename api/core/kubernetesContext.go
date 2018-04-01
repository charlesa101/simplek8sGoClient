package core

import (
	"context"
	"github.com/ericchiang/k8s"
)

type KubernetesContext struct {
	context.Context
	Config *k8s.Config
	Client *k8s.Client
}

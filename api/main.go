package api

import (
	"os"
	"flag"
	"path/filepath"
	"io/ioutil"
	"github.com/ericchiang/k8s"
	"github.com/ghodss/yaml"
	"context"
	"os/signal"
	"fmt"
	"github.com/k8sApi/api/server"
)

func main(){

	var kubeconfig *string

	if home := os.Getenv("HOME"); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	SetupHTTPServer(kubeconfig)

}

func SetupHTTPServer(kubeconfig *string) {

	data, err := ioutil.ReadFile(*kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	var config k8s.Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err.Error())
	}

	if err != nil {
		panic(err.Error())
	}

	client, err := k8s.NewClient(&config)
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	httpHandler := server.HttpHandler(ctx, client)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for _ = range c {
			ctx.Done()
			fmt.Println("Exiting Kubernetes Client!")
			os.Exit(0)
		}
	}()

	err = httpHandler.Bind(":9090")
	if err != nil {
		httpHandler.Done()
		panic(err.Error())
	}
}
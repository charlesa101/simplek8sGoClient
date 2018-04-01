
Simple Kubernetes GO API
========================

To Run this simple demo, make sure you have access to a running kubernetes cluster with `kubeconfig` accessible at
`$HOME/.kube/config`


0. Setup your `$GOPATH` with the correct toolchain access!

1. `go get` all needed dependencies

 -> `go get k8s.io/client-go/kubernetes`
 -> `go get github.com/gin-gonic/gin`
 -> `go get github.com/ghodss/yaml`
 -> `go get github.com/stretchr/testify/assert`

2. Run `go build github.com/k8sApi/api` to compile the API

3. Run `go install github.com/k8sApi/api` to install in your `$GOPATH/bin`

4. If steps 2. & 3. goes when then you should be able to just run `k8sApi` on the command line to start the kubernetes client webserver


The `k8sApi` web client serves the following endpoints

```
GET    /api/nodes
GET    /api/nodes/count
GET    /api/pods
GET    /api/pods/count
```


Example Request:

1. Get All pods in the cluster ->

REQUEST:
curl http://localhost:9090/api/pods/kube-system

RESPONSE:
![alt text](https://github.com/charlesa101/simplek8sGoClient/blob/master/docs/pods.png)

2. Get All pod count in the cluster ->

REQUEST:
curl http://localhost:9090/api/pods/kube-system/count

RESPONSE
`"{ 'pods' : 23 }"`



3. Get All nodes in the cluster ->

REQUEST:
curl http://localhost:9090/api/nodes

RESPONSE:
![alt text](https://github.com/charlesa101/simplek8sGoClient/blob/master/docs/nodes.png)

4. Get All node count in the cluster ->

REQUEST:
curl http://localhost:9090/api/nodes/count

RESPONSE
`"{ 'nodes' : 5 }"`

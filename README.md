# Hello Kubernetes

This project studies the provisioning of interconnected applications ("microservices") as Docker containers and their deployment to Kubernetes.
 
It has three subprojects, each containing a microservice:
* `backend-spring`: A trival "hello" Spring Boot application that echos the path part of its URL. For example, when called like http://localhost/World, it replies with "`[SPRING] Hello World`".
* `backend-golang`: Identical to `backend-spring`, but written in Go.
* `frontend-nodejs`: A node.js "frontend" service that forwards path argument of its URL to the "backend" services `backend-spring` and `backend-golang`, collects the results, and returns them.

## Preconditions
You need running Docker and Kubernetes installations, either on your workstation or at a public cloud provider.

For example, both Docker and Kubernetes come with [Docker Desktop](https://www.docker.com/products/docker-desktop) available for Mac and Windows.
An alternative is [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/), which also runs on Linux.

If you decide for a public cloud provider, you can use their managed Kubernetes offers, or setup a Kubernetes cluster with [Rancher](https://rancher.com/). 

## Build Docker Images
_Note_: You can omit this step if you are interested in Kubernetes only.
All images are available at [Docker Hub](https://hub.docker.com/) in repository [mouton4711/kubernetes](https://hub.docker.com/repository/docker/mouton4711/kubernetes).

```shell script
$ docker image build --tag mouton4711/kubernetes:backend-spring ./backend-spring
$ docker image build --tag mouton4711/kubernetes:backend-golang ./backend-golang
$ docker image build --tag mouton4711/kubernetes:frontend-nodejs ./frontend-nodejs
```

## Deploy to Kubernetes
```shell script
$ kubectl create configmap translation-config --from-literal=greetingLabel=Hello # Create config map from literals
$ kubectl apply -f ./backend-spring/kubernetes.yml
$ kubectl apply -f ./backend-golang/kubernetes.yml
$ kubectl apply -f ./frontend-nodejs/kubernetes.yml
```
Alternatively (and preferably!) you can combine all yaml files using the "kustomize" flag `-k` and deploy everything together:
```shell script
$ kubectl apply -k .
```

The services `backend-spring` and `backend-golang` do not expose their endpoints to the outside (they use the default service type `ClusterIP`).

Service `frontend-nodejs` has type [LoadBalancer](https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer).
If you run Docker Desktop, there is nothing additional to do, because Docker Deskop exposes services of type `LoadBalancer` to `localhost`.
You should be able to access the application from a web browser  
```
curl -s http://localhost/World
```
If your Kubernetes runs on a public cloud, you need to create an [Ingress](https://kubernetes.io/docs/concepts/services-networking/ingress/) service
that connects to the cloud provider's load balancer. Alternatively, you can setup an own load balancer like [ingress-nginx](https://github.com/kubernetes/ingress-nginx).
If you manage your Kubernetes cluster with [Rancher](https://rancher.com/), you can provision an ingress service using the [Rancher UI](https://rancher.com/docs/rancher/v2.x/en/k8s-in-rancher/load-balancers-and-ingress/ingress/).  

## Extras
Change greeting language:
```
$ kubectl edit configmap translation-config
# Change the value of "greeting" and save
$ kubectl rollout restart deployment backend-spring backend-golang frontend-nodejs
```

# Hello Kubernetes

This project studies the provisioning of interconnected applications ("microservices") as Docker containers and their deployment to Kubernetes.
 
It has four subprojects, each containing a microservice:
* `backend-spring`: A trival "hello" Spring Boot application that echos the path part of its URL. For example, when called like http://localhost/World, it replies with "`[SPRING] Hello World`".
* `backend-golang`: Identical to `backend-spring`, but written in Go.
* `backend-rust`: Identical to `backend-spring`, but written in Rust.
* `frontend-nodejs`: A node.js "frontend" service that forwards path argument of its URL to the "backend" services (`backend-spring`, `backend-golang`, `backend-rust`), collects the results, and returns them.

## Preconditions
You need running Docker and Kubernetes installations, either on your workstation or at a public cloud provider.

For example, both Docker and Kubernetes come with [Docker Desktop](https://www.docker.com/products/docker-desktop) available for Mac and Windows.
An alternative is [minikube](https://kubernetes.io/docs/setup/learning-environment/minikube/), which also runs on Linux.

If you decide for a public cloud provider, you can use their managed Kubernetes offers, or setup a Kubernetes cluster with [Rancher](https://rancher.com/). 

## Build Docker Images
_Note_: You can omit this step if you are interested in Kubernetes only.
All images are available at [Docker Hub](https://hub.docker.com/).

```shell script
$ sh build-local-images.sh
```

## Load Balancer Setup
The Ingress configuration expects a load balancer,
which is specified by the annotation `kubernetes.io/ingress.class` in file [ingress.yml](./ingress.yml).
It defaults to "nginx". To use nginx on Docker Desktop, please run
```shell
$ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
```

## TLS Configuration
Next, you need to set up TLS. This consists of two steps:
```shell
# 1. Create a self-signed certificate
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=localhost"
```

```shell
# 2. Add the certificate as secret to Kubernetes
kubectl create secret tls hello-kubernetes-tls --cert=tls.crt --key=tls.key
```
Make sure to use exactly this name, as the [ingress.yml](./ingress.yml) refers to it.

## Deploy to Kubernetes
```shell script
$ kubectl create configmap translation-config --from-literal=greetingLabel=Hello # Create config map from literals
$ kubectl apply -f ./backend-spring/kubernetes.yml
$ kubectl apply -f ./backend-golang/kubernetes.yml
$ kubectl apply -f ./backend-rust/kubernetes.yml
$ kubectl apply -f ./frontend-nodejs/kubernetes.yml
```
Alternatively (and preferably) you can combine all yaml files using the "kustomize" flag `-k` and deploy everything together:
```shell script
$ kubectl apply -k .
```

You should be able to access the application from a web browser (skip CA verification with flag `k`):  
```
$ curl -k https://localhost/World
```

If your Kubernetes runs on a public cloud, it depends on the cloud provider.
[GCP](https://console.cloud.google.com/) creates a Load Balancer if the `kubernetes.io/ingress.class` is "gce".
You can obtain the ephemeral IP address of the Ingress with
```
$ kubectl get ingress
```
and then use the displayed `ADDRESS` for calling the service: 
```
$ curl https://<ADDRESS>/World
```

## Extras
Change greeting language:
```
$ kubectl edit configmap translation-config
# Change the value of "greetingLabel" and save
$ kubectl rollout restart deployment backend-spring backend-golang backend-rust frontend-nodejs
```

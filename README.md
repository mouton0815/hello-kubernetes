# Hello Kubernetes

This project studies the deployment of an application set Kubernetes.
 
It has three subprojects:
* `backend-spring`: A trival "hello" Spring Boot application that echoes the path part of its URL. For example, when called like http://localhost:8080/World, it replies "Hello World".
* `backend-golang`: Identical to `backend-spring`, but written in Go.
* `frontend-nodejs`: A node.js "frontend" service that forwards path argument of its URL to the "backend" services `backend-spring` and `backend-golang`, collects the results, and returns them. 

## Preconditions
Running Docker and Kubernetes installations. For example both Docker and Kubernetes come with Docker Desktop.

## Spring Boot Application
```shell script
$ docker image build --tag backend-spring ./backend-spring
$ kubectl apply -f ./backend-spring/kubernetes.yml
```
Service `backend-spring` does not expose its endpoint to the load balancer.

## Golang Application
```shell script
$ docker image build --tag backend-golang ./backend-golang
$ kubectl apply -f ./backend-golang/kubernetes.yml
```
Service `backend-golang` does not expose its endpoint to the load balancer.

## Node Application
```shell script
$ docker image build --tag frontend-nodejs ./frontend-nodejs
$ kubectl apply -f ./frontend-nodejs/kubernetes.yml
```
http://localhost/World

## Deploy Everything Together
```shell script
$ kubectl kustomize . > kubernetes.yml 
$ kubectl apply -f ./kubernetes.yml
```

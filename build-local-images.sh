#!/bin/sh

# Build images for local testing
docker build -t mouton4711/hello_k8s_backend-rust ./backend-rust
docker build -t mouton4711/hello_k8s_backend-spring ./backend-spring
docker build -t mouton4711/hello_k8s_backend-golang ./backend-golang
docker build -t mouton4711/hello_k8s_frontend-nodejs ./frontend-nodejs

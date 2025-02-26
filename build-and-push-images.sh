#!/bin/sh

# Build and push multi-platform images
docker buildx build --platform linux/amd64,linux/arm64 --push -t mouton4711/hello_k8s_backend-rust ./backend-rust
docker buildx build --platform linux/amd64,linux/arm64 --push -t mouton4711/hello_k8s_backend-spring ./backend-spring
docker buildx build --platform linux/amd64,linux/arm64 --push -t mouton4711/hello_k8s_backend-golang ./backend-golang
docker buildx build --platform linux/amd64,linux/arm64 --push -t mouton4711/hello_k8s_frontend-nodejs ./frontend-nodejs

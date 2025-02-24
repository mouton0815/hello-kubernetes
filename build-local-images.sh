#!/bin/sh

# Build images for local testing
docker image build --tag mouton4711/hello_k8s_backend-rust ./backend-rust
docker image build --tag mouton4711/hello_k8s_backend-spring ./backend-spring
docker image build --tag mouton4711/hello_k8s_backend-golang ./backend-golang
docker image build --tag mouton4711/hello_k8s_frontend-nodejs ./frontend-nodejs

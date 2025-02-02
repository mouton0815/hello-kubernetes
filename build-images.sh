#!/bin/sh
docker image build --tag mouton4711/kubernetes:backend-rust ./backend-rust
docker image build --tag mouton4711/kubernetes:backend-spring ./backend-spring
docker image build --tag mouton4711/kubernetes:backend-golang ./backend-golang
docker image build --tag mouton4711/kubernetes:frontend-nodejs ./frontend-nodejs

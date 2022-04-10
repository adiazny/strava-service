SHELL := /bin/bash


# ==============================================================================
# Install dependencies

dev.setup.mac:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize

# ==============================================================================

run:
	go run cmd/services/strava-api/main.go

tidy:
	go mod tidy

down:
	kill -INT $(shell ps | grep exe/main | grep -v grep | sed -n 1,1p | cut -c1-5)

# ==============================================================================
# Building containers

# $(shell git rev-parse --short HEAD)
VERSION := 0.1.0

all: strava


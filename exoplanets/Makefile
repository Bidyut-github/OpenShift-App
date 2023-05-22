SHELL=/bin/bash

# version v1.0 is the legacy app version
# version v1.1.0 has "white background" and is based on ubi8 (check Containerfile)

VERSION ?= v1.1.0
REPO ?= quay.io/redhattraining

LOCAL_ADDRESS = $(shell hostname -I | awk '{print $$2}')

.PHONY: fmt build run run-without-db tag push pg pg-up pg-down

all: fmt build run

fmt:
	gofmt -w *.go

build:
	buildah build \
	  --build-arg VERSION=${VERSION} \
	  -t exoplanets:latest \
	  $(CURDIR) \
	;

rm:
	podman rm -f exoplanets

run:	rm
	podman run \
	  --name exoplanets \
	  -p 8080:8080 \
	  -e VERSION=${VERSION} \
	  -e DB_HOST=${LOCAL_ADDRESS} \
	  -e DB_PORT=5432 \
	  -e DB_USER=user \
	  -e DB_PASSWORD=password \
	  -e DB_NAME=postgres \
	  exoplanets:latest \
	;

run-without-db:	rm
	podman run \
	  --name exoplanets \
	  -p 8080:8080 \
	  exoplanets:latest \
	;

tag:
	# buildah tag exoplanets:latest ${REPO}/exoplanets:latest
	buildah tag exoplanets:latest ${REPO}/exoplanets:${VERSION}

push:
	# podman push ${REPO}/exoplanets:latest
	podman push ${REPO}/exoplanets:${VERSION}

POSTGRES_IMAGE=registry.redhat.io/rhel8/postgresql-13:1-118

pg:	pg-down pg-up

pg-up:
	podman run -d \
	  --name postgres \
	  -p ${LOCAL_ADDRESS}:5432:5432 \
	  -e POSTGRESQL_DATABASE=database \
	  -e POSTGRESQL_USER=user \
	  -e POSTGRESQL_PASSWORD=password \
	  -e POSTGRESQL_ADMIN_PASSWORD=postgres \
	  -e POSTGRESQL_MAX_CONNECTIONS=1000 \
	  ${POSTGRES_IMAGE} \
	;
	podman ps | grep postgres

pg-down:
	podman rm -f postgres

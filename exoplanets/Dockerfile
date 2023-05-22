#	Containerfile

# Be sure to match the 'builder' and 'app' image base version (both 'ubi8' or 'ubi9')

# About the 'version' LABEL
# version="v1.1.0" is based on 'ubi8'
# version="v1.1.1" is based on 'ubi9'

################################################################################
# -- Build stage --

# FROM	registry.redhat.io/ubi8/go-toolset:1.18.10-1	AS builder
FROM	registry.redhat.io/ubi9/go-toolset:1.18.10-4	AS builder

ENV	CGO_ENABLED=1

COPY	*.go .

RUN	go mod init exoplanets && \
	go mod tidy && \
	go get -u \
	  github.com/gorilla/mux@latest \
	  github.com/lib/pq@latest && \
	go build && \
	file exoplanets

################################################################################
# -- App stage --

# FROM	registry.redhat.io/ubi8/ubi-minimal:8.7-1107	AS app
FROM	registry.redhat.io/ubi9/ubi-minimal:9.1.0-1829	AS app

ARG	VERSION

LABEL	version=${VERSION}

WORKDIR	/srv

COPY	--from=builder --chmod=0755 /opt/app-root/src/exoplanets .
COPY	--chmod=0440 template.html .

RUN	chgrp -R 0 . && \
	chmod -R g=u .

USER	1001
EXPOSE	8080

CMD	["./exoplanets"]

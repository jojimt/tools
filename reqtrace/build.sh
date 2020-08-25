#!/bin/bash
set -e
: ${HUB_ID:=jojimt}

docker run --rm -m 16g -v ${GOPATH}:/go -w /go/src/jojimt/tools/reqtrace --network=host -e "CGO_ENABLED=0" -it noirolabs/gobuild1.14 go build -a -tags netgo -ldflags '-w -extldflags "-static"' .

docker build -t ${HUB_ID}/reqtrace -f Dockerfile .
docker push ${HUB_ID}/reqtrace

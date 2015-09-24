#!/bin/bash
echo "building statically-linked envtpl..."
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest sh -c 'CGO_ENABLED=0 go build -v -o bin/envtpl-linux64 ; CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o bin/envtpl-windows64.exe'

RELEASE_TAG=v0.2.0

case $1 in
"deploy")

go get github.com/aktau/github-release

git tag $RELEASE_TAG && git push --tags

github-release release --user pharmpress --repo envtpl --tag $RELEASE_TAG --name "envtpl$RELEASE_TAG" --description "my first release!" --pre-release

github-release upload --user pharmpress --repo envtpl --tag $RELEASE_TAG --name "envtpl-linux64" --file bin/envtpl-linux64

github-release upload --user pharmpress --repo envtpl --tag $RELEASE_TAG --name "envtpl-windows64.exe" --file bin/envtpl-windows64.exe

;;
esac


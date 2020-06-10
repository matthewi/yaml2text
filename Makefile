GIT_VER := $(shell git describe --tags)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
export GO111MODULE := on

.PHONY: test install clean dist release
cmd/yaml2text/yaml2text: *.go cmd/yaml2text/*.go
	cd cmd/yaml2text && go build -ldflags "-s -w -X main.version=${GIT_VER} -X main.buildDate=${DATE}" -gcflags="-trimpath=${PWD}"

install: cmd/yaml2text/yaml2text
	install cmd/yaml2text/yaml2text ${GOPATH}/bin

test:
	go test -race ./...

clean:
	rm -f cmd/yaml2text/yaml2text
	rm -fr dist/

dist:
	CGO_ENABLED=0 \
		goxz -pv=$(GIT_VER) \
		-build-ldflags="-s -w -X main.version=${GIT_VER} -X main.buildDate=${DATE}" \
		-os=darwin,linux -arch=amd64 -d=dist ./cmd/yaml2text

release:
	ghr -u mashiike -r yaml2text -n "$(GIT_VER)" $(GIT_VER) dist/

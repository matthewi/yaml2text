GIT_VER := $(shell git describe --tags)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
export GO111MODULE := on

.PHONY: test packages release clean

cmd/yaml2text/yaml2text: *.go cmd/yaml2text
	cd cmd/yaml2text && go build -ldflags "-s -w -X main.version=${GIT_VER} -X main.buildDate=${DATE}"

test:
	go test -v ./...

packages: *.go cmd/yaml2text
	cd cmd/yaml2text && gox -os="linux darwin" -arch="amd64" -output "../../pkg/{{.Dir}}-${GIT_VER}-{{.OS}}-{{.Arch}}" -ldflags "-s -w -X main.version=${GIT_VER} -X main.buildDate=${DATE}"
	cd pkg && find . -name "*${GIT_VER}*" -type f -exec zip {}.zip {} \;

release:
	ghr -u mashiike -r yaml2text -n "$(GIT_VER)" $(GIT_VER) pkg/

clean:
	rm -f cmd/yaml2text/yaml2text pkg/*

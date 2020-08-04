BINARY="target"
PACKAGES=`go list ./... | grep -v /vendor/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`
DATE=$(shell date '+%Y/%m/%d %H:%M:%S')

ifeq (${GOOS}, windows)
	BINARY=${BINARY}.exe
endif

build:
	go build -o ${BINARY} main.go


.PHONY: fmt
fmt:
	@gofmt -s -w ${GOFILES}

.PHONY: test
test:
	@go test -cpu=1,2,4 -v -tags integration ./...

.PHONY: docker
docker:
	@docker build -t hoojos/target:latest .

.PHONY: clean
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
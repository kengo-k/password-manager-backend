ROOT_DIR = $(shell pwd)
GOBIN = $(ROOT_DIR)/.bin
export PATH := $(ROOT_DIR)/.bin:$(PATH)

dev:
	cd src && go run ./main.go

test:
	cd src \
	&& go test -v -coverprofile=cover.out ./... \
	&& gcov2lcov -infile=cover.out -outfile=../cover.lcov \
	&& rm cover.out \
	&& cd .. && genhtml cover.lcov -o coverage_report \
	&& rm cover.lcov

install_gcov:
	GOBIN=$(GOBIN) go install github.com/jandelgado/gcov2lcov@v1.0.5

# ---- common ----

filter_file = gen|mock|thrift_gen|kitex_gen|clients|im_proto|kite.go|idls.go|\.pb\.go|\.gen\.go

all: clean tidy fmt lint test build

tidy:
	go mod tidy

fmt:
	find . -name '*.go' | grep -Ev "$(filter_file)" | xargs goimports -e -w -local "git.byted.org,code.byted.org"

.PHONY: build
build:
	bash ./build.sh

.PHONY: test
test:
	@bash .codebase/scripts/test.sh

cover:
	go tool cover -html=coverage.out

accurate_cover:
	@bash .codebase/scripts/test.sh accurate

lint:
	-golangci-lint run --new-from-rev=origin/master --timeout 2m0s

clean:
	rm -rf output *.out

generate:
	go generate ./...

upgrade:
	cat go.mod | grep -v '// indirect' | grep -v 'google.golang.org/protobuf' | grep -e '\t' | awk '{print $$1}' | xargs go get
	go mod tidy

# ---- customized ----

thrift:
	kitex -module code.byted.org/tiktok/im_conversation -service tiktok.im.conversation ../../im_cloud/idl_tiktok/server/im_cloud/backservice.thrift
	rm -rf ./kitex_gen
	sed -i '' -E -e 's#"code\.byted\.org/.*/kitex_gen/(.*)"#"code.byted.org/im_cloud/common_tiktok/gen/kitex_gen/\1"#g' main.go handler.go

#!/usr/bin/env bash
RUN_NAME="tiktok.im.lzz_test"

mkdir -p output/bin
cp script/* output/
chmod +x output/bootstrap.sh

if [ "$BUILD_TYPE" = "offline" -o "$BUILD_TYPE" = "test" -o "$BYTESUITE_USE_COVERAGE" = "1" ]; then
    go install code.byted.org/bet/go_coverage@tiktok_sg
    go_coverage annotate --extra-info=include:internal # use skip-files to skip files that not need calculate coverage, e.g. -skip-files=test*,example.go
fi

if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    go build -o output/bin/${RUN_NAME}
else
    go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi

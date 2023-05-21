#!/bin/bash
set -o pipefail
set -o xtrace

start=$(date +%s)

rm -rf *.out

export MallocNanoZone=0
export IS_UNIT_TEST_ENV=1
mode=$1
filter_pkg="\bgen\b|mock|kitex_gen|thrift_gen|clients|scripts"
filter_file="$filter_pkg|kite.go|idls.go|\.pb\.go|\.gen\.go|.*mock.*\.go"

if [ "x${mode}" == "x" ]; then
  go test -race -cover -coverprofile=coverage.out $(go list ./... | grep -Ev $filter_pkg) | sed -e 's/%.*/%/' -e 's/\[no test files\]/\[no test files\] ⚠️/'
elif [ "x${mode}" == "xgotestsum" ]; then
  gotestsum --junitfile=report.xml -- -race -cover -coverprofile=coverage.out $(go list ./... | grep -Ev $filter_pkg) -coverpkg $(go list ./... | grep -Ev $filter_pkg | tr "\n" "," | sed 's/.$//')
else
  go test -race -json -cover -coverprofile=coverage.out $(go list ./... | grep -Ev $filter_pkg) -coverpkg $(go list ./... | grep -Ev $filter_pkg | tr "\n" "," | sed 's/.$//') > report.json
fi
status=$?

set +o xtrace

[ $status -ne 0 ] && { printf -- "--------\nFailed\n"; exit $status; }
go tool cover -func=coverage.out | awk 'END {print $NF}'
printf -- "costs: %ss\n\n" $(($(date +%s) - $start))

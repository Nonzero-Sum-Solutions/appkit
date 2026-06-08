#!/bin/sh

test() {
  DATA_PATH="$(realpath ./test/data)"

  TEST_DATA_PATH=$DATA_PATH go test -test.v ./test/unit/...
}

case $1 in
  b  | build) go build .;;
  f  | format) go fmt ./*;;  
  h  | help)  echo "build|help|test|test_functional";;
  t  | test)  test;;
  tf | test_functional)  go test ./test/functional/...;;
  *) echo "Unknown argument";;
esac

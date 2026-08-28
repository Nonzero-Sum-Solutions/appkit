#!/bin/sh

test() {
  go test -test.v ./test/unit/...
}

gen() {
  designlanguage gen -d=false -t=go
}


case $1 in
  b  | build) go build .;;
  f  | format) go fmt ./*;;  
  g  | gen) gen;;
  h  | help)  echo "build|gen|help|test|test_functional";;
  t  | test)  test;;
  tf | test_functional)  go test ./test/functional/...;;
  *) echo "Unknown argument";;
esac

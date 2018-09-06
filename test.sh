#! /bin/bash

for i in $(find . -iname "*_test.go" -exec dirname {} \; | uniq | grep -v vendo)
do
   go test -race -cover $i;
done
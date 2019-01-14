#! /bin/bash

for i in $(find . -iname "*_test.go" -exec dirname {} \; | uniq | grep -v vendor | grep -v dependency)
do
   go test -race -cover $i;
done
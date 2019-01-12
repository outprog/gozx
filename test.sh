#! /bin/bash

for i in $(find . -iname "*_test.go" -exec dirname {} \; | uniq | grep -v vendor | grep -v dependence)
do
   go test -race -cover $i;
done
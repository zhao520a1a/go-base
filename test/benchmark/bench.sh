#!/usr/bin/env bash

pwd=$(pwd)
cd $pwd/
# 整体对比
#go test -benchmem -run=^$ -bench "^(BenchmarkString.*)$"

#go test -benchmem -run=none -count=20 -bench=BenchmarkStringCopy | tee old.txt
go test -benchmem -run=none -count=20 -bench=BenchmarkStringCopy | tee new.txt

#benchstat old.txt new.txt

cd $pwd
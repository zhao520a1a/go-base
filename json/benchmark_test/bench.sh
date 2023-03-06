#!/usr/bin/env bash

pwd=$(pwd)
export JSON_NO_ASYNC_GC=1

cd $pwd/
go test -benchmem -run=^$ -benchtime=1000000x -bench "^(BenchmarkEncoder_.*|BenchmarkDecoder_.*)$"

#go test -benchmem -run=^$ -benchtime=1000000x -bench "^(BenchmarkGet.*|BenchmarkSet.*)$"

#go test -benchmem -run=^$ -benchtime=10000x -bench "^(BenchmarkParser_.*)$"

unset JSON_NO_ASYNC_GC
cd $pwd
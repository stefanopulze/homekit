#!/usr/bin/env bash

target_os=$1
target_arch=$2
build_version=${3:-'0.0.1'}

echo "Compiling v${build_version} for os \"${target_os}:${target_arch}\""

env GOOS=$target_os GOARCH=$target_arch \
go build -ldflags="-X main.BuildVersion=0.1.0" -o ./output/homekit ./

echo "👍 Compile done"
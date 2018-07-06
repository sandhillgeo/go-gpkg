#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p $DIR/../bin

echo "******************"
echo "Formatting $DIR/../cmd/gpkg"
cd $DIR/../cmd/gpkg
go fmt
echo "Done formatting."
echo "******************"
echo "Building program for go-gpkg"
cd $DIR/../bin
####################################################
#echo "Building program for darwin"
#GOTAGS= CGO_ENABLED=1 GOOS=${GOOS} GOARCH=amd64 go build --tags "darwin" -o "gpkg_darwin_amd64" github.com/sandhillgeo/go-gpkg/cmd/gpkg
#if [[ "$?" != 0 ]] ; then
#    echo "Error building program for go-gpkg"
#    exit 1
#fi
#echo "Executable built at $(realpath $DIR/../bin/gpkg_darwin_amd64)"
####################################################
echo "Building program for linux"
GOTAGS= CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build --tags "linux" -o "gpkg_linux_amd64" github.com/sandhillgeo/go-gpkg/cmd/gpkg
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-gpkg"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin/gpkg_linux_amd64)"
####################################################
echo "Building program for Windows"
GOTAGS= CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc go build -o "gpkg_windows_amd64.exe" github.com/sandhillgeo/go-gpkg/cmd/gpkg
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-gpkg"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin/gpkg_windows_amd64.exe)"

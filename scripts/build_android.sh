#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p $DIR/../bin

echo "******************"
echo "Formatting $(realpath $DIR/../gpkg)"
cd $DIR/../gpkg
go fmt
echo "Done formatting."
echo "******************"
echo "Building AAR for gpkg"
cd $DIR/../bin
gomobile bind -target android github.com/sandhillgeo/go-gpkg/gpkg
if [[ "$?" != 0 ]] ; then
    echo "Error building program for gpkg"
    exit 1
fi
echo "Executable built at $(realpath $DIR/../bin)"

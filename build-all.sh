#!/bin/bash

#
# Script used to control the build
# 
# This script builds for 3 different Operating Systems:
# - Linux (i386)
# - OS X (i386)
# - Windows (i386)
#
# SPECIAL CONSIDERATIONS
# - Ensure the build environment is correctly configured for successful builds
#   - i.e.: use the "pre-build.sh" script to setup the build machine
#

BIN=df_gen
DESTDIR=bin

which go &>/dev/null
if [ $? -ne 0 -o -z "$GOROOT" ]
then
  echo "Failed to find GOROOT or the go executable - make sure it is set"
  exit 1
fi 

export GOARCH=386
export GOPATH=`pwd`

for d in win linux osx;
do
  mkdir -p $DESTDIR/$d
done

# Ensure formatting is followed 
go fmt ${BIN}

echo -e "\nBuilding for Linux(386)"
export GOOS=linux
go build -v -o $DESTDIR/linux/${BIN} ${BIN}

echo -e "\nBuilding for OS X/Darwin(386)"
export GOOS=darwin
go build -v -o $DESTDIR/osx/${BIN} ${BIN}

echo -e "\nBuilding for Windows(386)"
export GOOS=windows
go build -v -o $DESTDIR/win/${BIN}.exe ${BIN}

